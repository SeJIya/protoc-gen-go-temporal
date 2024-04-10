// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.10.5-next (38d49e8722013d965532492a3d6b9318a9a33971)
//	go go1.22.1
//	protoc (unknown)
//
// source: example/helloworld/v1/helloworld.proto
package helloworldv1

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	helpers "github.com/cludden/protoc-gen-go-temporal/pkg/helpers"
	scheme "github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
	gohomedir "github.com/mitchellh/go-homedir"
	v2 "github.com/urfave/cli/v2"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	testsuite "go.temporal.io/sdk/testsuite"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	protojson "google.golang.org/protobuf/encoding/protojson"
	"log/slog"
	"os"
	"sort"
)

// HelloWorldTaskQueue is the default task-queue for a example.helloworld.v1.HelloWorld worker
const HelloWorldTaskQueue = "hello-world"

// example.helloworld.v1.HelloWorld workflow names
const (
	HelloWorldWorkflowName = "HelloWorld"
)

// example.helloworld.v1.HelloWorld workflow id expressions
var (
	HelloWorldIdexpression = expression.MustParseExpression("hello_world/${! uuid_v4() }")
)

// example.helloworld.v1.HelloWorld activity names
const (
	HelloWorldActivityName = "HelloWorld"
)

// HelloWorldClient describes a client for a(n) example.helloworld.v1.HelloWorld worker
type HelloWorldClient interface {
	// HelloWorld describes a Temporal workflow and activity with the same name
	// and signature
	HelloWorld(ctx context.Context, req *HelloWorldInput, opts ...*HelloWorldOptions) (*HelloWorldOutput, error)

	// HelloWorldAsync starts a(n) HelloWorld workflow and returns a handle to the workflow run
	HelloWorldAsync(ctx context.Context, req *HelloWorldInput, opts ...*HelloWorldOptions) (HelloWorldRun, error)

	// GetHelloWorld retrieves a handle to an existing HelloWorld workflow execution
	GetHelloWorld(ctx context.Context, workflowID string, runID string) HelloWorldRun

	// CancelWorkflow requests cancellation of an existing workflow execution
	CancelWorkflow(ctx context.Context, workflowID string, runID string) error

	// TerminateWorkflow an existing workflow execution
	TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error
}

// helloWorldClient implements a temporal client for a example.helloworld.v1.HelloWorld service
type helloWorldClient struct {
	client client.Client
	log    *slog.Logger
}

// NewHelloWorldClient initializes a new example.helloworld.v1.HelloWorld client
func NewHelloWorldClient(c client.Client, options ...*helloWorldClientOptions) HelloWorldClient {
	var cfg *helloWorldClientOptions
	if len(options) > 0 {
		cfg = options[0]
	} else {
		cfg = NewHelloWorldClientOptions()
	}
	return &helloWorldClient{
		client: c,
		log:    cfg.getLogger(),
	}
}

// NewHelloWorldClientWithOptions initializes a new HelloWorld client with the given options
func NewHelloWorldClientWithOptions(c client.Client, opts client.Options, options ...*helloWorldClientOptions) (HelloWorldClient, error) {
	var err error
	c, err = client.NewClientFromExisting(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error initializing client with options: %w", err)
	}
	var cfg *helloWorldClientOptions
	if len(options) > 0 {
		cfg = options[0]
	} else {
		cfg = NewHelloWorldClientOptions()
	}
	return &helloWorldClient{
		client: c,
		log:    cfg.getLogger(),
	}, nil
}

// helloWorldClientOptions describes optional runtime configuration for a HelloWorldClient
type helloWorldClientOptions struct {
	log *slog.Logger
}

// NewHelloWorldClientOptions initializes a new helloWorldClientOptions value
func NewHelloWorldClientOptions() *helloWorldClientOptions {
	return &helloWorldClientOptions{}
}

// WithLogger can be used to override the default logger
func (opts *helloWorldClientOptions) WithLogger(l *slog.Logger) *helloWorldClientOptions {
	if l != nil {
		opts.log = l
	}
	return opts
}

// getLogger returns the configured logger, or the default logger
func (opts *helloWorldClientOptions) getLogger() *slog.Logger {
	if opts != nil && opts.log != nil {
		return opts.log
	}
	return slog.Default()
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
func (c *helloWorldClient) HelloWorld(ctx context.Context, req *HelloWorldInput, options ...*HelloWorldOptions) (*HelloWorldOutput, error) {
	run, err := c.HelloWorldAsync(ctx, req, options...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
func (c *helloWorldClient) HelloWorldAsync(ctx context.Context, req *HelloWorldInput, options ...*HelloWorldOptions) (HelloWorldRun, error) {
	opts := &client.StartWorkflowOptions{}
	if len(options) > 0 && options[0].opts != nil {
		opts = options[0].opts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = HelloWorldTaskQueue
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(HelloWorldIdexpression, req.ProtoReflect())
		if err != nil {
			return nil, fmt.Errorf("error evaluating id expression for \"HelloWorld\" workflow: %w", err)
		}
		opts.ID = id
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, HelloWorldWorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &helloWorldRun{
		client: c,
		run:    run,
	}, nil
}

// GetHelloWorld fetches an existing HelloWorld execution
func (c *helloWorldClient) GetHelloWorld(ctx context.Context, workflowID string, runID string) HelloWorldRun {
	return &helloWorldRun{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}
}

// CancelWorkflow requests cancellation of an existing workflow execution
func (c *helloWorldClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return c.client.CancelWorkflow(ctx, workflowID, runID)
}

// TerminateWorkflow terminates an existing workflow execution
func (c *helloWorldClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	return c.client.TerminateWorkflow(ctx, workflowID, runID, reason, details...)
}

// HelloWorldOptions provides configuration for a HelloWorld workflow operation
type HelloWorldOptions struct {
	opts *client.StartWorkflowOptions
}

// NewHelloWorldOptions initializes a new HelloWorldOptions value
func NewHelloWorldOptions() *HelloWorldOptions {
	return &HelloWorldOptions{}
}

// WithStartWorkflowOptions sets the initial client.StartWorkflowOptions
func (opts *HelloWorldOptions) WithStartWorkflowOptions(options client.StartWorkflowOptions) *HelloWorldOptions {
	opts.opts = &options
	return opts
}

// HelloWorldRun describes a(n) HelloWorld workflow run
type HelloWorldRun interface {
	// ID returns the workflow ID
	ID() string

	// RunID returns the workflow instance ID
	RunID() string

	// Run returns the inner client.WorkflowRun
	Run() client.WorkflowRun

	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) (*HelloWorldOutput, error)

	// Cancel requests cancellation of a workflow in execution, returning an error if applicable
	Cancel(ctx context.Context) error

	// Terminate terminates a workflow in execution, returning an error if applicable
	Terminate(ctx context.Context, reason string, details ...interface{}) error
}

// helloWorldRun provides an internal implementation of a(n) HelloWorldRunRun
type helloWorldRun struct {
	client *helloWorldClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *helloWorldRun) ID() string {
	return r.run.GetID()
}

// Run returns the inner client.WorkflowRun
func (r *helloWorldRun) Run() client.WorkflowRun {
	return r.run
}

// RunID returns the execution ID
func (r *helloWorldRun) RunID() string {
	return r.run.GetRunID()
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *helloWorldRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *helloWorldRun) Get(ctx context.Context) (*HelloWorldOutput, error) {
	var resp HelloWorldOutput
	if err := r.run.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *helloWorldRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// Reference to generated workflow functions
var (
	// HelloWorld describes a Temporal workflow and activity with the same name
	// and signature
	HelloWorldFunction func(workflow.Context, *HelloWorldInput) (*HelloWorldOutput, error)
)

// HelloWorldWorkflows provides methods for initializing new example.helloworld.v1.HelloWorld workflow values
type HelloWorldWorkflows interface {
	// HelloWorld describes a Temporal workflow and activity with the same name
	// and signature
	HelloWorld(ctx workflow.Context, input *HelloWorldWorkflowInput) (HelloWorldWorkflow, error)
}

// RegisterHelloWorldWorkflows registers example.helloworld.v1.HelloWorld workflows with the given worker
func RegisterHelloWorldWorkflows(r worker.WorkflowRegistry, workflows HelloWorldWorkflows) {
	RegisterHelloWorldWorkflow(r, workflows.HelloWorld)
}

// RegisterHelloWorldWorkflow registers a example.helloworld.v1.HelloWorld.HelloWorld workflow with the given worker
func RegisterHelloWorldWorkflow(r worker.WorkflowRegistry, wf func(workflow.Context, *HelloWorldWorkflowInput) (HelloWorldWorkflow, error)) {
	HelloWorldFunction = buildHelloWorld(wf)
	r.RegisterWorkflowWithOptions(HelloWorldFunction, workflow.RegisterOptions{Name: HelloWorldWorkflowName})
}

// buildHelloWorld converts a HelloWorld workflow struct into a valid workflow function
func buildHelloWorld(ctor func(workflow.Context, *HelloWorldWorkflowInput) (HelloWorldWorkflow, error)) func(workflow.Context, *HelloWorldInput) (*HelloWorldOutput, error) {
	return func(ctx workflow.Context, req *HelloWorldInput) (*HelloWorldOutput, error) {
		input := &HelloWorldWorkflowInput{
			Req: req,
		}
		wf, err := ctor(ctx, input)
		if err != nil {
			return nil, err
		}
		if initializable, ok := wf.(helpers.Initializable); ok {
			if err := initializable.Initialize(ctx); err != nil {
				return nil, err
			}
		}
		return wf.Execute(ctx)
	}
}

// HelloWorldWorkflowInput describes the input to a(n) HelloWorld workflow constructor
type HelloWorldWorkflowInput struct {
	Req *HelloWorldInput
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
//
// workflow details: (id: "hello_world/${! uuid_v4() }")
type HelloWorldWorkflow interface {
	// Execute defines the entrypoint to a(n) HelloWorld workflow
	Execute(ctx workflow.Context) (*HelloWorldOutput, error)
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
func HelloWorldChild(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldChildOptions) (*HelloWorldOutput, error) {
	childRun, err := HelloWorldChildAsync(ctx, req, options...)
	if err != nil {
		return nil, err
	}
	return childRun.Get(ctx)
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
func HelloWorldChildAsync(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldChildOptions) (*HelloWorldChildRun, error) {
	opts := &workflow.ChildWorkflowOptions{}
	if len(options) > 0 && options[0].opts != nil {
		opts = options[0].opts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = HelloWorldTaskQueue
	}
	if opts.WorkflowID == "" {
		id, err := expression.EvalExpression(HelloWorldIdexpression, req.ProtoReflect())
		if err != nil {
			panic(err)
		}
		opts.WorkflowID = id
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return &HelloWorldChildRun{Future: workflow.ExecuteChildWorkflow(ctx, HelloWorldWorkflowName, req)}, nil
}

// HelloWorldChildOptions provides configuration for a HelloWorld workflow operation
type HelloWorldChildOptions struct {
	opts *workflow.ChildWorkflowOptions
}

// NewHelloWorldChildOptions initializes a new HelloWorldChildOptions value
func NewHelloWorldChildOptions() *HelloWorldChildOptions {
	return &HelloWorldChildOptions{}
}

// WithChildWorkflowOptions sets the initial client.StartWorkflowOptions
func (opts *HelloWorldChildOptions) WithChildWorkflowOptions(options workflow.ChildWorkflowOptions) *HelloWorldChildOptions {
	opts.opts = &options
	return opts
}

// HelloWorldChildRun describes a child HelloWorld workflow run
type HelloWorldChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *HelloWorldChildRun) Get(ctx workflow.Context) (*HelloWorldOutput, error) {
	var resp HelloWorldOutput
	if err := r.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *HelloWorldChildRun) Select(sel workflow.Selector, fn func(*HelloWorldChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *HelloWorldChildRun) SelectStart(sel workflow.Selector, fn func(*HelloWorldChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *HelloWorldChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// HelloWorldActivities describes available worker activities
type HelloWorldActivities interface {
	// HelloWorld describes a Temporal workflow and activity with the same name
	// and signature
	HelloWorld(ctx context.Context, req *HelloWorldInput) (*HelloWorldOutput, error)
}

// RegisterHelloWorldActivities registers activities with a worker
func RegisterHelloWorldActivities(r worker.ActivityRegistry, activities HelloWorldActivities) {
	RegisterHelloWorldActivity(r, activities.HelloWorld)
}

// RegisterHelloWorldActivity registers a HelloWorld activity
func RegisterHelloWorldActivity(r worker.ActivityRegistry, fn func(context.Context, *HelloWorldInput) (*HelloWorldOutput, error)) {
	r.RegisterActivityWithOptions(fn, activity.RegisterOptions{
		Name: HelloWorldActivityName,
	})
}

// HelloWorldFuture describes a(n) HelloWorld activity execution
type HelloWorldFuture struct {
	Future workflow.Future
}

// Get blocks on the activity's completion, returning the response
func (f *HelloWorldFuture) Get(ctx workflow.Context) (*HelloWorldOutput, error) {
	var resp HelloWorldOutput
	if err := f.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds the activity's completion to the selector, callback can be nil
func (f *HelloWorldFuture) Select(sel workflow.Selector, fn func(*HelloWorldFuture)) workflow.Selector {
	return sel.AddFuture(f.Future, func(workflow.Future) {
		if fn != nil {
			fn(f)
		}
	})
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
func HelloWorld(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldActivityOptions) (*HelloWorldOutput, error) {
	var opts *HelloWorldActivityOptions
	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	} else {
		opts = NewHelloWorldActivityOptions()
	}
	if opts.opts == nil {
		opts.opts = &workflow.ActivityOptions{}
	}
	if opts.opts.TaskQueue == "" {
		opts.opts.TaskQueue = HelloWorldTaskQueue
	}
	if opts.opts.StartToCloseTimeout == 0 {
		opts.opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithActivityOptions(ctx, *opts.opts)
	var activity any
	activity = HelloWorldActivityName
	future := &HelloWorldFuture{Future: workflow.ExecuteActivity(ctx, activity, req)}
	return future.Get(ctx)
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
func HelloWorldAsync(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldActivityOptions) *HelloWorldFuture {
	var opts *HelloWorldActivityOptions
	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	} else {
		opts = NewHelloWorldActivityOptions()
	}
	if opts.opts == nil {
		opts.opts = &workflow.ActivityOptions{}
	}
	if opts.opts.TaskQueue == "" {
		opts.opts.TaskQueue = HelloWorldTaskQueue
	}
	if opts.opts.StartToCloseTimeout == 0 {
		opts.opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithActivityOptions(ctx, *opts.opts)
	var activity any
	activity = HelloWorldActivityName
	future := &HelloWorldFuture{Future: workflow.ExecuteActivity(ctx, activity, req)}
	return future
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
func HelloWorldLocal(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldLocalActivityOptions) (*HelloWorldOutput, error) {
	var opts *HelloWorldLocalActivityOptions
	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	} else {
		opts = NewHelloWorldLocalActivityOptions()
	}
	if opts.opts == nil {
		opts.opts = &workflow.LocalActivityOptions{}
	}
	if opts.opts.StartToCloseTimeout == 0 {
		opts.opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts.opts)
	var activity any
	if opts.fn != nil {
		activity = opts.fn
	} else {
		activity = HelloWorldActivityName
	}
	future := &HelloWorldFuture{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
	return future.Get(ctx)
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
func HelloWorldLocalAsync(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldLocalActivityOptions) *HelloWorldFuture {
	var opts *HelloWorldLocalActivityOptions
	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	} else {
		opts = NewHelloWorldLocalActivityOptions()
	}
	if opts.opts == nil {
		opts.opts = &workflow.LocalActivityOptions{}
	}
	if opts.opts.StartToCloseTimeout == 0 {
		opts.opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts.opts)
	var activity any
	if opts.fn != nil {
		activity = opts.fn
	} else {
		activity = HelloWorldActivityName
	}
	future := &HelloWorldFuture{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
	return future
}

// HelloWorldLocalActivityOptions provides configuration for a local HelloWorld activity
type HelloWorldLocalActivityOptions struct {
	fn   func(context.Context, *HelloWorldInput) (*HelloWorldOutput, error)
	opts *workflow.LocalActivityOptions
}

// NewHelloWorldLocalActivityOptions sets default LocalActivityOptions
func NewHelloWorldLocalActivityOptions() *HelloWorldLocalActivityOptions {
	return &HelloWorldLocalActivityOptions{}
}

// Local provides a local HelloWorld activity implementation
func (opts *HelloWorldLocalActivityOptions) Local(fn func(context.Context, *HelloWorldInput) (*HelloWorldOutput, error)) *HelloWorldLocalActivityOptions {
	opts.fn = fn
	return opts
}

// WithLocalActivityOptions sets default LocalActivityOptions
func (opts *HelloWorldLocalActivityOptions) WithLocalActivityOptions(options workflow.LocalActivityOptions) *HelloWorldLocalActivityOptions {
	opts.opts = &options
	return opts
}

// HelloWorldActivityOptions provides configuration for a(n) HelloWorld activity
type HelloWorldActivityOptions struct {
	opts *workflow.ActivityOptions
}

// NewHelloWorldActivityOptions sets default ActivityOptions
func NewHelloWorldActivityOptions() *HelloWorldActivityOptions {
	return &HelloWorldActivityOptions{}
}

// WithActivityOptions sets default ActivityOptions
func (opts *HelloWorldActivityOptions) WithActivityOptions(options workflow.ActivityOptions) *HelloWorldActivityOptions {
	opts.opts = &options
	return opts
}

// TestClient provides a testsuite-compatible Client
type TestHelloWorldClient struct {
	env       *testsuite.TestWorkflowEnvironment
	workflows HelloWorldWorkflows
}

var _ HelloWorldClient = &TestHelloWorldClient{}

// NewTestHelloWorldClient initializes a new TestHelloWorldClient value
func NewTestHelloWorldClient(env *testsuite.TestWorkflowEnvironment, workflows HelloWorldWorkflows, activities HelloWorldActivities) *TestHelloWorldClient {
	if workflows != nil {
		RegisterHelloWorldWorkflows(env, workflows)
	}
	if activities != nil {
		RegisterHelloWorldActivities(env, activities)
	}
	return &TestHelloWorldClient{env, workflows}
}

// HelloWorld executes a(n) HelloWorld workflow in the test environment
func (c *TestHelloWorldClient) HelloWorld(ctx context.Context, req *HelloWorldInput, opts ...*HelloWorldOptions) (*HelloWorldOutput, error) {
	run, err := c.HelloWorldAsync(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// HelloWorldAsync executes a(n) HelloWorld workflow in the test environment
func (c *TestHelloWorldClient) HelloWorldAsync(ctx context.Context, req *HelloWorldInput, options ...*HelloWorldOptions) (HelloWorldRun, error) {
	opts := &client.StartWorkflowOptions{}
	if len(options) > 0 && options[0].opts != nil {
		opts = options[0].opts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = HelloWorldTaskQueue
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(HelloWorldIdexpression, req.ProtoReflect())
		if err != nil {
			return nil, fmt.Errorf("error evaluating id expression for \"HelloWorld\" workflow: %w", err)
		}
		opts.ID = id
	}
	return &testHelloWorldRun{client: c, env: c.env, opts: opts, req: req, workflows: c.workflows}, nil
}

// GetHelloWorld is a noop
func (c *TestHelloWorldClient) GetHelloWorld(ctx context.Context, workflowID string, runID string) HelloWorldRun {
	return &testHelloWorldRun{env: c.env, workflows: c.workflows}
}

// CancelWorkflow requests cancellation of an existing workflow execution
func (c *TestHelloWorldClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	c.env.CancelWorkflow()
	return nil
}

// TerminateWorkflow terminates an existing workflow execution
func (c *TestHelloWorldClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	return c.CancelWorkflow(ctx, workflowID, runID)
}

var _ HelloWorldRun = &testHelloWorldRun{}

// testHelloWorldRun provides convenience methods for interacting with a(n) HelloWorld workflow in the test environment
type testHelloWorldRun struct {
	client    *TestHelloWorldClient
	env       *testsuite.TestWorkflowEnvironment
	opts      *client.StartWorkflowOptions
	req       *HelloWorldInput
	workflows HelloWorldWorkflows
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *testHelloWorldRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get retrieves a test HelloWorld workflow result
func (r *testHelloWorldRun) Get(context.Context) (*HelloWorldOutput, error) {
	r.env.ExecuteWorkflow(HelloWorldWorkflowName, r.req)
	if !r.env.IsWorkflowCompleted() {
		return nil, errors.New("workflow in progress")
	}
	if err := r.env.GetWorkflowError(); err != nil {
		return nil, err
	}
	var result HelloWorldOutput
	if err := r.env.GetWorkflowResult(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ID returns a test HelloWorld workflow run's workflow ID
func (r *testHelloWorldRun) ID() string {
	if r.opts != nil {
		return r.opts.ID
	}
	return ""
}

// Run noop implementation
func (r *testHelloWorldRun) Run() client.WorkflowRun {
	return nil
}

// RunID noop implementation
func (r *testHelloWorldRun) RunID() string {
	return ""
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *testHelloWorldRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// HelloWorldCliOptions describes runtime configuration for example.helloworld.v1.HelloWorld cli
type HelloWorldCliOptions struct {
	after            func(*v2.Context) error
	before           func(*v2.Context) error
	clientForCommand func(*v2.Context) (client.Client, error)
	worker           func(*v2.Context, client.Client) (worker.Worker, error)
}

// NewHelloWorldCliOptions initializes a new HelloWorldCliOptions value
func NewHelloWorldCliOptions() *HelloWorldCliOptions {
	return &HelloWorldCliOptions{}
}

// WithAfter injects a custom After hook to be run after any command invocation
func (opts *HelloWorldCliOptions) WithAfter(fn func(*v2.Context) error) *HelloWorldCliOptions {
	opts.after = fn
	return opts
}

// WithBefore injects a custom Before hook to be run prior to any command invocation
func (opts *HelloWorldCliOptions) WithBefore(fn func(*v2.Context) error) *HelloWorldCliOptions {
	opts.before = fn
	return opts
}

// WithClient provides a Temporal client factory for use by commands
func (opts *HelloWorldCliOptions) WithClient(fn func(*v2.Context) (client.Client, error)) *HelloWorldCliOptions {
	opts.clientForCommand = fn
	return opts
}

// WithWorker provides an method for initializing a worker
func (opts *HelloWorldCliOptions) WithWorker(fn func(*v2.Context, client.Client) (worker.Worker, error)) *HelloWorldCliOptions {
	opts.worker = fn
	return opts
}

// NewHelloWorldCli initializes a cli for a(n) example.helloworld.v1.HelloWorld service
func NewHelloWorldCli(options ...*HelloWorldCliOptions) (*v2.App, error) {
	commands, err := newHelloWorldCommands(options...)
	if err != nil {
		return nil, fmt.Errorf("error initializing subcommands: %w", err)
	}
	return &v2.App{
		Name:     "hello-world",
		Commands: commands,
	}, nil
}

// NewHelloWorldCliCommand initializes a cli command for a example.helloworld.v1.HelloWorld service with subcommands for each query, signal, update, and workflow
func NewHelloWorldCliCommand(options ...*HelloWorldCliOptions) (*v2.Command, error) {
	subcommands, err := newHelloWorldCommands(options...)
	if err != nil {
		return nil, fmt.Errorf("error initializing subcommands: %w", err)
	}
	return &v2.Command{
		Name:        "hello-world",
		Subcommands: subcommands,
	}, nil
}

// newHelloWorldCommands initializes (sub)commands for a example.helloworld.v1.HelloWorld cli or command
func newHelloWorldCommands(options ...*HelloWorldCliOptions) ([]*v2.Command, error) {
	opts := &HelloWorldCliOptions{}
	if len(options) > 0 {
		opts = options[0]
	}
	if opts.clientForCommand == nil {
		opts.clientForCommand = func(*v2.Context) (client.Client, error) {
			return client.Dial(client.Options{})
		}
	}
	commands := []*v2.Command{
		{
			Name:                   "hello-world",
			Usage:                  "HelloWorld describes a Temporal workflow and activity with the same name and signature",
			Category:               "WORKFLOWS",
			UseShortOptionHandling: true,
			Before:                 opts.before,
			After:                  opts.after,
			Flags: []v2.Flag{
				&v2.BoolFlag{
					Name:    "detach",
					Usage:   "run workflow in the background and print workflow and execution id",
					Aliases: []string{"d"},
				},
				&v2.StringFlag{
					Name:    "task-queue",
					Usage:   "task queue name",
					Aliases: []string{"t"},
					EnvVars: []string{"TEMPORAL_TASK_QUEUE_NAME", "TEMPORAL_TASK_QUEUE", "TASK_QUEUE_NAME", "TASK_QUEUE"},
					Value:   "hello-world",
				},
				&v2.StringFlag{
					Name:    "input-file",
					Usage:   "path to json-formatted input file",
					Aliases: []string{"f"},
				},
				&v2.StringFlag{
					Name:     "name",
					Usage:    "set the value of the operation's \"Name\" parameter",
					Category: "INPUT",
				},
			},
			Action: func(cmd *v2.Context) error {
				tc, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer tc.Close()
				c := NewHelloWorldClient(tc)
				req, err := UnmarshalCliFlagsToHelloWorldInput(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				opts := client.StartWorkflowOptions{}
				if tq := cmd.String("task-queue"); tq != "" {
					opts.TaskQueue = tq
				}
				run, err := c.HelloWorldAsync(cmd.Context, req, NewHelloWorldOptions().WithStartWorkflowOptions(opts))
				if err != nil {
					return fmt.Errorf("error starting %s workflow: %w", HelloWorldWorkflowName, err)
				}
				if cmd.Bool("detach") {
					fmt.Println("success")
					fmt.Printf("workflow id: %s\n", run.ID())
					fmt.Printf("run id: %s\n", run.RunID())
					return nil
				}
				if resp, err := run.Get(cmd.Context); err != nil {
					return err
				} else {
					b, err := protojson.Marshal(resp)
					if err != nil {
						return fmt.Errorf("error serializing response json: %w", err)
					}
					var out bytes.Buffer
					if err := json.Indent(&out, b, "", "  "); err != nil {
						return fmt.Errorf("error formatting json: %w", err)
					}
					fmt.Println(out.String())
					return nil
				}
			},
		},
	}
	if opts.worker != nil {
		commands = append(commands, []*v2.Command{
			{
				Name:                   "worker",
				Usage:                  "runs a example.helloworld.v1.HelloWorld worker process",
				UseShortOptionHandling: true,
				Before:                 opts.before,
				After:                  opts.after,
				Action: func(cmd *v2.Context) error {
					c, err := opts.clientForCommand(cmd)
					if err != nil {
						return fmt.Errorf("error initializing client for command: %w", err)
					}
					defer c.Close()
					w, err := opts.worker(cmd, c)
					if opts.worker != nil {
						if err != nil {
							return fmt.Errorf("error initializing worker: %w", err)
						}
					}
					if err := w.Start(); err != nil {
						return fmt.Errorf("error starting worker: %w", err)
					}
					defer w.Stop()
					<-cmd.Context.Done()
					return nil
				},
			},
		}...)
	}
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Name < commands[j].Name
	})
	return commands, nil
}

// UnmarshalCliFlagsToHelloWorldInput unmarshals a HelloWorldInput from command line flags
func UnmarshalCliFlagsToHelloWorldInput(cmd *v2.Context) (*HelloWorldInput, error) {
	var result HelloWorldInput
	var hasValues bool
	if cmd.IsSet("input-file") {
		inputFile, err := gohomedir.Expand(cmd.String("input-file"))
		if err != nil {
			inputFile = cmd.String("input-file")
		}
		b, err := os.ReadFile(inputFile)
		if err != nil {
			return nil, fmt.Errorf("error reading input-file: %w", err)
		}
		if err := protojson.Unmarshal(b, &result); err != nil {
			return nil, fmt.Errorf("error parsing input-file json: %w", err)
		}
		hasValues = true
	}
	if cmd.IsSet("name") {
		hasValues = true
		result.Name = cmd.String("name")
	}
	if !hasValues {
		return nil, nil
	}
	return &result, nil
}

// WithHelloWorldSchemeTypes registers all HelloWorld protobuf types with the given scheme
func WithHelloWorldSchemeTypes() scheme.Option {
	return func(s *scheme.Scheme) {
		s.RegisterType(File_example_helloworld_v1_helloworld_proto.Messages().ByName("HelloWorldInput"))
		s.RegisterType(File_example_helloworld_v1_helloworld_proto.Messages().ByName("HelloWorldOutput"))
	}
}
