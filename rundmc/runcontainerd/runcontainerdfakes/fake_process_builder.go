// Code generated by counterfeiter. DO NOT EDIT.
package runcontainerdfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	"code.cloudfoundry.org/guardian/rundmc/runcontainerd"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeProcessBuilder struct {
	BuildProcessStub        func(bndl goci.Bndl, spec garden.ProcessSpec, uid, gid int) *specs.Process
	buildProcessMutex       sync.RWMutex
	buildProcessArgsForCall []struct {
		bndl goci.Bndl
		spec garden.ProcessSpec
		uid  int
		gid  int
	}
	buildProcessReturns struct {
		result1 *specs.Process
	}
	buildProcessReturnsOnCall map[int]struct {
		result1 *specs.Process
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProcessBuilder) BuildProcess(bndl goci.Bndl, spec garden.ProcessSpec, uid int, gid int) *specs.Process {
	fake.buildProcessMutex.Lock()
	ret, specificReturn := fake.buildProcessReturnsOnCall[len(fake.buildProcessArgsForCall)]
	fake.buildProcessArgsForCall = append(fake.buildProcessArgsForCall, struct {
		bndl goci.Bndl
		spec garden.ProcessSpec
		uid  int
		gid  int
	}{bndl, spec, uid, gid})
	fake.recordInvocation("BuildProcess", []interface{}{bndl, spec, uid, gid})
	fake.buildProcessMutex.Unlock()
	if fake.BuildProcessStub != nil {
		return fake.BuildProcessStub(bndl, spec, uid, gid)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.buildProcessReturns.result1
}

func (fake *FakeProcessBuilder) BuildProcessCallCount() int {
	fake.buildProcessMutex.RLock()
	defer fake.buildProcessMutex.RUnlock()
	return len(fake.buildProcessArgsForCall)
}

func (fake *FakeProcessBuilder) BuildProcessArgsForCall(i int) (goci.Bndl, garden.ProcessSpec, int, int) {
	fake.buildProcessMutex.RLock()
	defer fake.buildProcessMutex.RUnlock()
	return fake.buildProcessArgsForCall[i].bndl, fake.buildProcessArgsForCall[i].spec, fake.buildProcessArgsForCall[i].uid, fake.buildProcessArgsForCall[i].gid
}

func (fake *FakeProcessBuilder) BuildProcessReturns(result1 *specs.Process) {
	fake.BuildProcessStub = nil
	fake.buildProcessReturns = struct {
		result1 *specs.Process
	}{result1}
}

func (fake *FakeProcessBuilder) BuildProcessReturnsOnCall(i int, result1 *specs.Process) {
	fake.BuildProcessStub = nil
	if fake.buildProcessReturnsOnCall == nil {
		fake.buildProcessReturnsOnCall = make(map[int]struct {
			result1 *specs.Process
		})
	}
	fake.buildProcessReturnsOnCall[i] = struct {
		result1 *specs.Process
	}{result1}
}

func (fake *FakeProcessBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildProcessMutex.RLock()
	defer fake.buildProcessMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProcessBuilder) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ runcontainerd.ProcessBuilder = new(FakeProcessBuilder)