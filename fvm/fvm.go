package fvm

import (
	"github.com/onflow/cadence/runtime"

	"github.com/dapperlabs/flow-go/model/flow"
)

type Invokable interface {
	Parse(ctx Context, ledger Ledger) (Invokable, error)
	Invoke(ctx Context, ledger Ledger) (*InvocationResult, error)
}

// A VirtualMachine augments the Cadence runtime with Flow host functionality.
type VirtualMachine struct {
	rt    runtime.Runtime
	chain flow.Chain
}

// New creates a new virtual machine instance with the provided runtime.
func New(rt runtime.Runtime, chain flow.Chain) *VirtualMachine {
	return &VirtualMachine{
		rt:    rt,
		chain: chain,
	}
}

// NewContext initializes a new execution context with the provided options.
func (vm *VirtualMachine) NewContext(opts ...Option) Context {
	return newContext(vm.rt, defaultOptions(vm.chain), opts...)
}