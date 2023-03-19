// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 0.3.1-next (c2718b2010cd3901e803cdf5557e6b377022ab36)
//	go go1.19.6
//	protoc (unknown)
//
// source: example.proto
package foov1

import (
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	v1 "go.temporal.io/api/enums/v1"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
)

// Foo workflow names
const (
	LockAccountWorkflowName = "mycompany.foo.v1.Foo.LockAccountWorkflow"
	TransferWorkflowName    = "mycompany.foo.v1.Foo.TransferWorkflow"
)

// Foo id prefixes
const (
	LockAccountIDPrefix = "lock"
	TransferIDPrefix    = "transfer"
)

// Foo signal names
const (
	AcquireLeaseSignalName  = "mycompany.foo.v1.Foo.AcquireLeaseSignal"
	LeaseAcquiredSignalName = "mycompany.foo.v1.Foo.LeaseAcquiredSignal"
	RenewLeaseSignalName    = "mycompany.foo.v1.Foo.RenewLeaseSignal"
	RevokeLeaseSignalName   = "mycompany.foo.v1.Foo.RevokeLeaseSignal"
)

// Foo activity names
const (
	DepositActivityName  = "mycompany.foo.v1.Foo.DepositActivity"
	WithdrawActivityName = "mycompany.foo.v1.Foo.WithdrawActivity"
)

// Client describes a client for a Foo worker
type Client interface {
	// LockAccount executes a LockAccount workflow and blocks until error or response received
	LockAccount(ctx context.Context, opts *client.StartWorkflowOptions, req *LockAccountRequest) error
	// ExecuteLockAccount executes a LockAccount workflow
	ExecuteLockAccount(ctx context.Context, opts *client.StartWorkflowOptions, req *LockAccountRequest) (LockAccountRun, error)
	// GetLockAccount retrieves a LockAccount workflow execution
	GetLockAccount(ctx context.Context, workflowID string, runID string) (LockAccountRun, error)
	// StartLockAccountWithAcquireLease sends a AcquireLease signal to a LockAccount workflow, starting it if not present
	StartLockAccountWithAcquireLease(ctx context.Context, opts *client.StartWorkflowOptions, req *LockAccountRequest, signal *AcquireLeaseSignal) (LockAccountRun, error)
	// Transfer executes a Transfer workflow and blocks until error or response received
	Transfer(ctx context.Context, opts *client.StartWorkflowOptions, req *TransferRequest) (*TransferResponse, error)
	// ExecuteTransfer executes a Transfer workflow
	ExecuteTransfer(ctx context.Context, opts *client.StartWorkflowOptions, req *TransferRequest) (TransferRun, error)
	// GetTransfer retrieves a Transfer workflow execution
	GetTransfer(ctx context.Context, workflowID string, runID string) (TransferRun, error)
	// AcquireLease sends a AcquireLease signal to an existing workflow
	AcquireLease(ctx context.Context, workflowID string, runID string, signal *AcquireLeaseSignal) error
	// LeaseAcquired sends a LeaseAcquired signal to an existing workflow
	LeaseAcquired(ctx context.Context, workflowID string, runID string, signal *LeaseAcquiredSignal) error
	// RenewLease sends a RenewLease signal to an existing workflow
	RenewLease(ctx context.Context, workflowID string, runID string, signal *RenewLeaseSignal) error
	// RevokeLease sends a RevokeLease signal to an existing workflow
	RevokeLease(ctx context.Context, workflowID string, runID string, signal *RevokeLeaseSignal) error
}

// Compile-time check that workflowClient satisfies Client
var _ Client = &workflowClient{}

// workflowClient implements a temporal client for a Foo service
type workflowClient struct {
	client client.Client
}

// NewClient initializes a new Foo client
func NewClient(c client.Client) Client {
	return &workflowClient{client: c}
}

// LockAccount executes a LockAccount workflow and blocks until error or response received
func (c *workflowClient) LockAccount(ctx context.Context, opts *client.StartWorkflowOptions, req *LockAccountRequest) error {
	run, err := c.ExecuteLockAccount(ctx, opts, req)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// ExecuteLockAccount starts a LockAccount workflow
func (c *workflowClient) ExecuteLockAccount(ctx context.Context, opts *client.StartWorkflowOptions, req *LockAccountRequest) (LockAccountRun, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "foo-v1"
	}
	if opts.ID == "" {
		opts.ID = fmt.Sprintf("%s/%v", LockAccountIDPrefix, req.GetAccount())
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, LockAccountWorkflowName, req)
	if run == nil || err != nil {
		return nil, err
	}
	return &lockAccountRun{
		client: c,
		run:    run,
	}, nil
}

// GetLockAccount fetches an existing LockAccount execution
func (c *workflowClient) GetLockAccount(ctx context.Context, workflowID string, runID string) (LockAccountRun, error) {
	return &lockAccountRun{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}, nil
}

// StartLockAccountWithAcquireLease starts a LockAccount workflow and sends a AcquireLease signal in a transaction
func (c *workflowClient) StartLockAccountWithAcquireLease(ctx context.Context, opts *client.StartWorkflowOptions, req *LockAccountRequest, signal *AcquireLeaseSignal) (LockAccountRun, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "foo-v1"
	}
	if opts.ID == "" {
		opts.ID = fmt.Sprintf("%s/%v", LockAccountIDPrefix, req.GetAccount())
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	run, err := c.client.SignalWithStartWorkflow(ctx, opts.ID, AcquireLeaseSignalName, signal, *opts, LockAccountWorkflowName, req)
	if run == nil || err != nil {
		return nil, err
	}
	return &lockAccountRun{
		client: c,
		run:    run,
	}, nil
}

// Transfer executes a Transfer workflow and blocks until error or response received
func (c *workflowClient) Transfer(ctx context.Context, opts *client.StartWorkflowOptions, req *TransferRequest) (*TransferResponse, error) {
	run, err := c.ExecuteTransfer(ctx, opts, req)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// ExecuteTransfer starts a Transfer workflow
func (c *workflowClient) ExecuteTransfer(ctx context.Context, opts *client.StartWorkflowOptions, req *TransferRequest) (TransferRun, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "foo-v1"
	}
	if opts.ID == "" {
		opts.ID = fmt.Sprintf("%s/%v/%v/%s", TransferIDPrefix, req.GetSrc(), req.GetDest(), uuid.New().String())
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, TransferWorkflowName, req)
	if run == nil || err != nil {
		return nil, err
	}
	return &transferRun{
		client: c,
		run:    run,
	}, nil
}

// GetTransfer fetches an existing Transfer execution
func (c *workflowClient) GetTransfer(ctx context.Context, workflowID string, runID string) (TransferRun, error) {
	return &transferRun{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}, nil
}

// AcquireLease sends a AcquireLease signal to an existing workflow
func (c *workflowClient) AcquireLease(ctx context.Context, workflowID string, runID string, signal *AcquireLeaseSignal) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, AcquireLeaseSignalName, signal)
}

// LeaseAcquired sends a LeaseAcquired signal to an existing workflow
func (c *workflowClient) LeaseAcquired(ctx context.Context, workflowID string, runID string, signal *LeaseAcquiredSignal) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, LeaseAcquiredSignalName, signal)
}

// RenewLease sends a RenewLease signal to an existing workflow
func (c *workflowClient) RenewLease(ctx context.Context, workflowID string, runID string, signal *RenewLeaseSignal) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, RenewLeaseSignalName, signal)
}

// RevokeLease sends a RevokeLease signal to an existing workflow
func (c *workflowClient) RevokeLease(ctx context.Context, workflowID string, runID string, signal *RevokeLeaseSignal) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, RevokeLeaseSignalName, signal)
}

// LockAccountRun describes a LockAccount workflow run
type LockAccountRun interface {
	// ID returns the workflow ID
	ID() string
	// RunID returns the workflow instance ID
	RunID() string
	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) error
	// AcquireLease sends a AcquireLease signal to the workflow
	AcquireLease(ctx context.Context, req *AcquireLeaseSignal) error
	// RenewLease sends a RenewLease signal to the workflow
	RenewLease(ctx context.Context, req *RenewLeaseSignal) error
	// RevokeLease sends a RevokeLease signal to the workflow
	RevokeLease(ctx context.Context, req *RevokeLeaseSignal) error
}

// lockAccountRun provides an internal implementation of a LockAccountRun
type lockAccountRun struct {
	client *workflowClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *lockAccountRun) ID() string {
	return r.run.GetID()
}

// RunID returns the execution ID
func (r *lockAccountRun) RunID() string {
	return r.run.GetRunID()
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *lockAccountRun) Get(ctx context.Context) error {
	return r.run.Get(ctx, nil)
}

// AcquireLease sends a AcquireLease signal to the workflow
func (r *lockAccountRun) AcquireLease(ctx context.Context, req *AcquireLeaseSignal) error {
	return r.client.AcquireLease(ctx, r.ID(), "", req)
}

// RenewLease sends a RenewLease signal to the workflow
func (r *lockAccountRun) RenewLease(ctx context.Context, req *RenewLeaseSignal) error {
	return r.client.RenewLease(ctx, r.ID(), "", req)
}

// RevokeLease sends a RevokeLease signal to the workflow
func (r *lockAccountRun) RevokeLease(ctx context.Context, req *RevokeLeaseSignal) error {
	return r.client.RevokeLease(ctx, r.ID(), "", req)
}

// TransferRun describes a Transfer workflow run
type TransferRun interface {
	// ID returns the workflow ID
	ID() string
	// RunID returns the workflow instance ID
	RunID() string
	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) (*TransferResponse, error)
	// LeaseAcquired sends a LeaseAcquired signal to the workflow
	LeaseAcquired(ctx context.Context, req *LeaseAcquiredSignal) error
}

// transferRun provides an internal implementation of a TransferRun
type transferRun struct {
	client *workflowClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *transferRun) ID() string {
	return r.run.GetID()
}

// RunID returns the execution ID
func (r *transferRun) RunID() string {
	return r.run.GetRunID()
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *transferRun) Get(ctx context.Context) (*TransferResponse, error) {
	var resp TransferResponse
	if err := r.run.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LeaseAcquired sends a LeaseAcquired signal to the workflow
func (r *transferRun) LeaseAcquired(ctx context.Context, req *LeaseAcquiredSignal) error {
	return r.client.LeaseAcquired(ctx, r.ID(), "", req)
}

// Workflows provides methods for initializing new Foo workflow values
type Workflows interface {
	// LockAccount initializes a new LockAccountWorkflow value
	LockAccount(ctx workflow.Context, input *LockAccountInput) (LockAccountWorkflow, error)
	// Transfer initializes a new TransferWorkflow value
	Transfer(ctx workflow.Context, input *TransferInput) (TransferWorkflow, error)
}

// RegisterWorkflows registers Foo workflows with the given worker
func RegisterWorkflows(r worker.Registry, workflows Workflows) {
	RegisterLockAccountWorkflow(r, workflows.LockAccount)
	RegisterTransferWorkflow(r, workflows.Transfer)
}

// RegisterLockAccountWorkflow registers a LockAccount workflow with the given worker
func RegisterLockAccountWorkflow(r worker.Registry, wf func(workflow.Context, *LockAccountInput) (LockAccountWorkflow, error)) {
	r.RegisterWorkflowWithOptions(buildLockAccount(wf), workflow.RegisterOptions{Name: LockAccountWorkflowName})
}

// buildLockAccount converts a LockAccount workflow struct into a valid workflow function
func buildLockAccount(wf func(workflow.Context, *LockAccountInput) (LockAccountWorkflow, error)) func(workflow.Context, *LockAccountRequest) error {
	return (&lockAccount{wf}).LockAccount
}

// lockAccount provides an LockAccount method for calling the user's implementation
type lockAccount struct {
	ctor func(workflow.Context, *LockAccountInput) (LockAccountWorkflow, error)
}

// LockAccount constructs a new LockAccount value and executes it
func (w *lockAccount) LockAccount(ctx workflow.Context, req *LockAccountRequest) error {
	input := &LockAccountInput{
		Req: req,
		AcquireLease: &AcquireLease{
			Channel: workflow.GetSignalChannel(ctx, AcquireLeaseSignalName),
		},
		RenewLease: &RenewLease{
			Channel: workflow.GetSignalChannel(ctx, RenewLeaseSignalName),
		},
		RevokeLease: &RevokeLease{
			Channel: workflow.GetSignalChannel(ctx, RevokeLeaseSignalName),
		},
	}
	wf, err := w.ctor(ctx, input)
	if err != nil {
		return err
	}
	return wf.Execute(ctx)
}

// LockAccountInput describes the input to a LockAccount workflow constructor
type LockAccountInput struct {
	Req          *LockAccountRequest
	AcquireLease *AcquireLease
	RenewLease   *RenewLease
	RevokeLease  *RevokeLease
}

// LockAccountWorkflow describes a LockAccount workflow implementation
type LockAccountWorkflow interface {
	// Execute a LockAccount workflow
	Execute(ctx workflow.Context) error
}

// LockAccountChild executes a child LockAccount workflow
func LockAccountChild(ctx workflow.Context, opts *workflow.ChildWorkflowOptions, req *LockAccountRequest) LockAccountChildRun {
	if opts == nil {
		childOpts := workflow.GetChildWorkflowOptions(ctx)
		opts = &childOpts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "foo-v1"
	}
	if opts.WorkflowID == "" {
		opts.WorkflowID = fmt.Sprintf("%s/%v", LockAccountIDPrefix, req.GetAccount())
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return LockAccountChildRun{
		Future: workflow.ExecuteChildWorkflow(ctx, "LockAccountWorkflowName", req),
	}
}

// LockAccountChildRun describes a child LockAccount workflow run
type LockAccountChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *LockAccountChildRun) Get(ctx workflow.Context) error {
	if err := r.Future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *LockAccountChildRun) Select(sel workflow.Selector, fn func(LockAccountChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *LockAccountChildRun) SelectStart(sel workflow.Selector, fn func(LockAccountChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *LockAccountChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// AcquireLease sends the corresponding signal request to the child workflow
func (r *LockAccountChildRun) AcquireLease(ctx workflow.Context, input *AcquireLeaseSignal) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, AcquireLeaseSignalName, input)
}

// RenewLease sends the corresponding signal request to the child workflow
func (r *LockAccountChildRun) RenewLease(ctx workflow.Context, input *RenewLeaseSignal) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, RenewLeaseSignalName, input)
}

// RevokeLease sends the corresponding signal request to the child workflow
func (r *LockAccountChildRun) RevokeLease(ctx workflow.Context, input *RevokeLeaseSignal) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, RevokeLeaseSignalName, input)
}

// RegisterTransferWorkflow registers a Transfer workflow with the given worker
func RegisterTransferWorkflow(r worker.Registry, wf func(workflow.Context, *TransferInput) (TransferWorkflow, error)) {
	r.RegisterWorkflowWithOptions(buildTransfer(wf), workflow.RegisterOptions{Name: TransferWorkflowName})
}

// buildTransfer converts a Transfer workflow struct into a valid workflow function
func buildTransfer(wf func(workflow.Context, *TransferInput) (TransferWorkflow, error)) func(workflow.Context, *TransferRequest) (*TransferResponse, error) {
	return (&transfer{wf}).Transfer
}

// transfer provides an Transfer method for calling the user's implementation
type transfer struct {
	ctor func(workflow.Context, *TransferInput) (TransferWorkflow, error)
}

// Transfer constructs a new Transfer value and executes it
func (w *transfer) Transfer(ctx workflow.Context, req *TransferRequest) (*TransferResponse, error) {
	input := &TransferInput{
		Req: req,
		LeaseAcquired: &LeaseAcquired{
			Channel: workflow.GetSignalChannel(ctx, LeaseAcquiredSignalName),
		},
	}
	wf, err := w.ctor(ctx, input)
	if err != nil {
		return nil, err
	}
	return wf.Execute(ctx)
}

// TransferInput describes the input to a Transfer workflow constructor
type TransferInput struct {
	Req           *TransferRequest
	LeaseAcquired *LeaseAcquired
}

// TransferWorkflow describes a Transfer workflow implementation
type TransferWorkflow interface {
	// Execute a Transfer workflow
	Execute(ctx workflow.Context) (*TransferResponse, error)
}

// TransferChild executes a child Transfer workflow
func TransferChild(ctx workflow.Context, opts *workflow.ChildWorkflowOptions, req *TransferRequest) TransferChildRun {
	if opts == nil {
		childOpts := workflow.GetChildWorkflowOptions(ctx)
		opts = &childOpts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "foo-v1"
	}
	if opts.WorkflowID == "" {
		opts.WorkflowID = fmt.Sprintf("%s/%v/%v/%s", TransferIDPrefix, req.GetSrc(), req.GetDest(), uuid.New().String())
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return TransferChildRun{
		Future: workflow.ExecuteChildWorkflow(ctx, "TransferWorkflowName", req),
	}
}

// TransferChildRun describes a child Transfer workflow run
type TransferChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *TransferChildRun) Get(ctx workflow.Context) (*TransferResponse, error) {
	var resp TransferResponse
	if err := r.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *TransferChildRun) Select(sel workflow.Selector, fn func(TransferChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *TransferChildRun) SelectStart(sel workflow.Selector, fn func(TransferChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *TransferChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// LeaseAcquired sends the corresponding signal request to the child workflow
func (r *TransferChildRun) LeaseAcquired(ctx workflow.Context, input *LeaseAcquiredSignal) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, LeaseAcquiredSignalName, input)
}

// AcquireLease describes a AcquireLease signal
type AcquireLease struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a AcquireLease signal is received
func (s *AcquireLease) Receive(ctx workflow.Context) (*AcquireLeaseSignal, bool) {
	var resp AcquireLeaseSignal
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a AcquireLease signal without blocking
func (s *AcquireLease) ReceiveAsync() *AcquireLeaseSignal {
	var resp AcquireLeaseSignal
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// Select checks for a AcquireLease signal without blocking
func (s *AcquireLease) Select(sel workflow.Selector, fn func(*AcquireLeaseSignal)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// AcquireLeaseExternal sends a AcquireLease signal to an existing workflow
func AcquireLeaseExternal(ctx workflow.Context, workflowID string, runID string, req *AcquireLeaseSignal) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, AcquireLeaseSignalName, req)
}

// LeaseAcquired describes a LeaseAcquired signal
type LeaseAcquired struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a LeaseAcquired signal is received
func (s *LeaseAcquired) Receive(ctx workflow.Context) (*LeaseAcquiredSignal, bool) {
	var resp LeaseAcquiredSignal
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a LeaseAcquired signal without blocking
func (s *LeaseAcquired) ReceiveAsync() *LeaseAcquiredSignal {
	var resp LeaseAcquiredSignal
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// Select checks for a LeaseAcquired signal without blocking
func (s *LeaseAcquired) Select(sel workflow.Selector, fn func(*LeaseAcquiredSignal)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// LeaseAcquiredExternal sends a LeaseAcquired signal to an existing workflow
func LeaseAcquiredExternal(ctx workflow.Context, workflowID string, runID string, req *LeaseAcquiredSignal) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, LeaseAcquiredSignalName, req)
}

// RenewLease describes a RenewLease signal
type RenewLease struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a RenewLease signal is received
func (s *RenewLease) Receive(ctx workflow.Context) (*RenewLeaseSignal, bool) {
	var resp RenewLeaseSignal
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a RenewLease signal without blocking
func (s *RenewLease) ReceiveAsync() *RenewLeaseSignal {
	var resp RenewLeaseSignal
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// Select checks for a RenewLease signal without blocking
func (s *RenewLease) Select(sel workflow.Selector, fn func(*RenewLeaseSignal)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// RenewLeaseExternal sends a RenewLease signal to an existing workflow
func RenewLeaseExternal(ctx workflow.Context, workflowID string, runID string, req *RenewLeaseSignal) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, RenewLeaseSignalName, req)
}

// RevokeLease describes a RevokeLease signal
type RevokeLease struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a RevokeLease signal is received
func (s *RevokeLease) Receive(ctx workflow.Context) (*RevokeLeaseSignal, bool) {
	var resp RevokeLeaseSignal
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a RevokeLease signal without blocking
func (s *RevokeLease) ReceiveAsync() *RevokeLeaseSignal {
	var resp RevokeLeaseSignal
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// Select checks for a RevokeLease signal without blocking
func (s *RevokeLease) Select(sel workflow.Selector, fn func(*RevokeLeaseSignal)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// RevokeLeaseExternal sends a RevokeLease signal to an existing workflow
func RevokeLeaseExternal(ctx workflow.Context, workflowID string, runID string, req *RevokeLeaseSignal) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, RevokeLeaseSignalName, req)
}

// Activities describes available worker activites
type Activities interface {
	// Deposit amount into an account
	Deposit(ctx context.Context, req *DepositRequest) (*DepositResponse, error)
	// Withdraw amount from an account
	Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error)
}

// RegisterActivities registers activities with a worker
func RegisterActivities(r worker.Registry, activities Activities) {
	RegisterDepositActivity(r, activities.Deposit)
	RegisterWithdrawActivity(r, activities.Withdraw)
}

// RegisterDepositActivity registers a Deposit activity
func RegisterDepositActivity(r worker.Registry, fn func(context.Context, *DepositRequest) (*DepositResponse, error)) {
	r.RegisterActivityWithOptions(fn, activity.RegisterOptions{
		Name: DepositActivityName,
	})
}

// DepositFuture describes a Deposit activity execution
type DepositFuture struct {
	Future workflow.Future
}

// Get blocks on a Deposit execution, returning the response
func (f *DepositFuture) Get(ctx workflow.Context) (*DepositResponse, error) {
	var resp DepositResponse
	if err := f.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds the Deposit completion to the selector, callback can be nil
func (f *DepositFuture) Select(sel workflow.Selector, fn func(*DepositFuture)) workflow.Selector {
	return sel.AddFuture(f.Future, func(workflow.Future) {
		if fn != nil {
			fn(f)
		}
	})
}

// Deposit amount into an account
func Deposit(ctx workflow.Context, opts *workflow.ActivityOptions, req *DepositRequest) *DepositFuture {
	if opts == nil {
		activityOpts := workflow.GetActivityOptions(ctx)
		opts = &activityOpts
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{
			MaximumAttempts: int32(5),
		}
	}
	if opts.ScheduleToCloseTimeout == 0 {
		opts.ScheduleToCloseTimeout = 120000000000 // 2m0s
	}
	ctx = workflow.WithActivityOptions(ctx, *opts)
	return &DepositFuture{Future: workflow.ExecuteActivity(ctx, DepositActivityName, req)}
}

// Deposit amount into an account
func DepositLocal(ctx workflow.Context, opts *workflow.LocalActivityOptions, fn func(context.Context, *DepositRequest) (*DepositResponse, error), req *DepositRequest) *DepositFuture {
	if opts == nil {
		activityOpts := workflow.GetLocalActivityOptions(ctx)
		opts = &activityOpts
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{
			MaximumAttempts: int32(5),
		}
	}
	if opts.ScheduleToCloseTimeout == 0 {
		opts.ScheduleToCloseTimeout = 120000000000 // 2m0s
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts)
	var activity any
	if fn == nil {
		activity = DepositActivityName
	} else {
		activity = fn
	}
	return &DepositFuture{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
}

// RegisterWithdrawActivity registers a Withdraw activity
func RegisterWithdrawActivity(r worker.Registry, fn func(context.Context, *WithdrawRequest) (*WithdrawResponse, error)) {
	r.RegisterActivityWithOptions(fn, activity.RegisterOptions{
		Name: WithdrawActivityName,
	})
}

// WithdrawFuture describes a Withdraw activity execution
type WithdrawFuture struct {
	Future workflow.Future
}

// Get blocks on a Withdraw execution, returning the response
func (f *WithdrawFuture) Get(ctx workflow.Context) (*WithdrawResponse, error) {
	var resp WithdrawResponse
	if err := f.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds the Withdraw completion to the selector, callback can be nil
func (f *WithdrawFuture) Select(sel workflow.Selector, fn func(*WithdrawFuture)) workflow.Selector {
	return sel.AddFuture(f.Future, func(workflow.Future) {
		if fn != nil {
			fn(f)
		}
	})
}

// Withdraw amount from an account
func Withdraw(ctx workflow.Context, opts *workflow.ActivityOptions, req *WithdrawRequest) *WithdrawFuture {
	if opts == nil {
		activityOpts := workflow.GetActivityOptions(ctx)
		opts = &activityOpts
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{
			MaximumAttempts: int32(5),
		}
	}
	if opts.ScheduleToCloseTimeout == 0 {
		opts.ScheduleToCloseTimeout = 120000000000 // 2m0s
	}
	ctx = workflow.WithActivityOptions(ctx, *opts)
	return &WithdrawFuture{Future: workflow.ExecuteActivity(ctx, WithdrawActivityName, req)}
}

// Withdraw amount from an account
func WithdrawLocal(ctx workflow.Context, opts *workflow.LocalActivityOptions, fn func(context.Context, *WithdrawRequest) (*WithdrawResponse, error), req *WithdrawRequest) *WithdrawFuture {
	if opts == nil {
		activityOpts := workflow.GetLocalActivityOptions(ctx)
		opts = &activityOpts
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{
			MaximumAttempts: int32(5),
		}
	}
	if opts.ScheduleToCloseTimeout == 0 {
		opts.ScheduleToCloseTimeout = 120000000000 // 2m0s
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts)
	var activity any
	if fn == nil {
		activity = WithdrawActivityName
	} else {
		activity = fn
	}
	return &WithdrawFuture{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
}
