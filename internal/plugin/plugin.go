package plugin

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	g "github.com/dave/jennifer/jen"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	deprecatedComment = "Deprecated: Do not use."
)

type Config struct {
	CliCategories              bool
	CliEnabled                 bool
	DisableWorkflowInputRename bool
	DocsOut                    string
	DocsTemplate               string
	EnableCodec                bool
	EnablePatchSupport         bool
	EnableXNS                  bool
	WorkflowUpdateEnabled      bool
}

// Plugin provides a protoc plugin for generating temporal workers and clients in go
type Plugin struct {
	*protogen.Plugin

	Commit  string
	Version string
	cfg     *Config
	flags   *pflag.FlagSet
}

func New(commit, version string) *Plugin {
	var cfg Config

	flags := pflag.NewFlagSet("plugin", pflag.ExitOnError)
	flags.BoolVar(&cfg.CliEnabled, "cli-enabled", false, "enable cli generation")
	flags.BoolVar(&cfg.CliCategories, "cli-categories", true, "enable cli categories")
	flags.BoolVar(&cfg.DisableWorkflowInputRename, "disable-workflow-input-rename", false, "disable renaming of \"<Workflow>WorkflowInput\"")
	flags.StringVar(&cfg.DocsOut, "docs-out", "", "docs output path")
	flags.StringVar(&cfg.DocsTemplate, "docs-template", "basic", "built-in template name or path to custom template file")
	flags.BoolVar(&cfg.EnableCodec, "enable-codec", false, "enables experimental codec support")
	flags.BoolVar(&cfg.EnablePatchSupport, "enable-patch-support", false, "enables support for alta/protopatch renaming")
	flags.BoolVar(&cfg.EnableXNS, "enable-xns", false, "enable experimental cross-namespace workflow client")
	flags.BoolVar(&cfg.WorkflowUpdateEnabled, "workflow-update-enabled", false, "enable experimental workflow update")

	return &Plugin{
		Commit:  commit,
		Version: version,
		cfg:     &cfg,
		flags:   flags,
	}
}

// Param provides a protogen ParamFunc handler
func (p *Plugin) Param(key, value string) error {
	return p.flags.Set(key, value)
}

// Run defines the plugin entrypoint
func (p *Plugin) Run(plugin *protogen.Plugin) error {
	p.Plugin = plugin

	for _, file := range p.Files {
		if !file.Generate {
			continue
		}

		pkgName := string(file.GoPackageName)
		f := g.NewFile(string(pkgName))
		genCodeGenerationHeader(p, f, file)

		var xns *g.File
		var xnsGoPackageName, xnsFilePath string
		var hasXNS bool
		if p.cfg.EnableXNS {
			xnsGoPackageName = fmt.Sprintf("%sxns", file.GoPackageName)
			xns = g.NewFile(xnsGoPackageName)
			genCodeGenerationHeader(p, xns, file)

			prefixToSlash := filepath.ToSlash(file.GeneratedFilenamePrefix)
			xnsFilePath = path.Join(
				path.Dir(prefixToSlash),
				xnsGoPackageName,
				path.Base(prefixToSlash),
			)
		}

		var hasContent bool
		for _, service := range file.Services {
			svc, err := parseService(plugin, p.cfg, file, service)
			if err != nil {
				return fmt.Errorf("error parsing service %s: %w", service.GoName, err)
			}
			if len(svc.activities) == 0 && len(svc.workflows) == 0 && len(svc.signals) == 0 && len(svc.queries) == 0 {
				continue
			}

			svc.render(f)
			svc.renderTestClient(f)
			if svc.cfg.CliEnabled {
				svc.renderCLI(f)
			}
			if svc.cfg.EnableXNS {
				svc.renderXNS(xns)
				hasXNS = true
			}
			if svc.cfg.EnableCodec {
				svc.renderCodec(f)
			}
			hasContent = true
		}

		if !hasContent {
			continue
		}

		if err := f.Render(p.NewGeneratedFile(fmt.Sprintf("%s_temporal.pb.go", file.GeneratedFilenamePrefix), file.GoImportPath)); err != nil {
			return fmt.Errorf("error rendering file: %w", err)
		}
		if hasXNS {
			if err := xns.Render(p.NewGeneratedFile(
				fmt.Sprintf("%s_xns_temporal.pb.go", xnsFilePath),
				protogen.GoImportPath(path.Join(
					string(file.GoImportPath),
					xnsGoPackageName,
				)),
			)); err != nil {
				return fmt.Errorf("error rendering file: %w", err)
			}
		}
	}

	if p.cfg.DocsOut != "" {
		if err := renderDocs(p.Plugin, p.cfg); err != nil {
			fmt.Fprintf(os.Stderr, "error rendering docs: %v", err)
		}
	}
	return nil
}

func commentf[T any, PT interface {
	*T
	Comment(string) *g.Statement
	Commentf(string, ...any) *g.Statement
}](c PT, methods []*protogen.Method, defaultMsg string, a ...any) {
	var deprecated bool
	for _, method := range methods {
		deprecated = isDeprecated(method)
		if deprecated {
			break
		}
	}
	if len(methods) == 1 && methods[0].Comments.Leading.String() != "" {
		comment := strings.TrimSuffix(methods[0].Comments.Leading.String(), "\n")
		c.Comment(comment)
		if deprecated && !strings.Contains(comment, "Deprecated:") {
			c.Comment(" ")
			c.Comment(deprecatedComment)
		}
	} else {
		c.Commentf(defaultMsg, a...)
		if deprecated {
			c.Comment(" ")
			c.Comment(deprecatedComment)
		}
	}
}

func genCodeGenerationHeader(p *Plugin, f *g.File, target *protogen.File) {
	f.PackageComment("Code generated by protoc-gen-go_temporal. DO NOT EDIT.")
	f.PackageComment("versions: ")
	f.PackageComment(fmt.Sprintf("    protoc-gen-go_temporal %s (%s)", p.Version, p.Commit))
	f.PackageComment(fmt.Sprintf("    go %s", runtime.Version()))
	compilerVersion := p.Plugin.Request.CompilerVersion
	if compilerVersion != nil {
		f.PackageComment(fmt.Sprintf("    protoc %s", compilerVersion.String()))
	} else {
		f.PackageComment("    protoc (unknown)")
	}

	f.PackageComment(fmt.Sprintf("source: %s", target.Desc.Path()))
}

func isDeprecated(method *protogen.Method) bool {
	return method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated()
}

// isEmpty checks if the message is a google.protobuf.Empty message
func isEmpty(m *protogen.Message) bool {
	return m.Desc.FullName() == "google.protobuf.Empty"
}

func methodSet(methods ...*protogen.Method) []*protogen.Method {
	return methods
}
