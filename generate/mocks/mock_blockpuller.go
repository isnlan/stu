// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"stu/generate"
	"sync"
)

type FakeBlockPuller struct {
	PullBlockStub        func(uint64) error
	pullBlockMutex       sync.RWMutex
	pullBlockArgsForCall []struct {
		arg1 uint64
	}
	pullBlockReturns struct {
		result1 error
	}
	pullBlockReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlockPuller) PullBlock(arg1 uint64) error {
	fake.pullBlockMutex.Lock()
	ret, specificReturn := fake.pullBlockReturnsOnCall[len(fake.pullBlockArgsForCall)]
	fake.pullBlockArgsForCall = append(fake.pullBlockArgsForCall, struct {
		arg1 uint64
	}{arg1})
	fake.recordInvocation("PullBlock", []interface{}{arg1})
	fake.pullBlockMutex.Unlock()
	if fake.PullBlockStub != nil {
		return fake.PullBlockStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pullBlockReturns
	return fakeReturns.result1
}

func (fake *FakeBlockPuller) PullBlockCallCount() int {
	fake.pullBlockMutex.RLock()
	defer fake.pullBlockMutex.RUnlock()
	return len(fake.pullBlockArgsForCall)
}

func (fake *FakeBlockPuller) PullBlockCalls(stub func(uint64) error) {
	fake.pullBlockMutex.Lock()
	defer fake.pullBlockMutex.Unlock()
	fake.PullBlockStub = stub
}

func (fake *FakeBlockPuller) PullBlockArgsForCall(i int) uint64 {
	fake.pullBlockMutex.RLock()
	defer fake.pullBlockMutex.RUnlock()
	argsForCall := fake.pullBlockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockPuller) PullBlockReturns(result1 error) {
	fake.pullBlockMutex.Lock()
	defer fake.pullBlockMutex.Unlock()
	fake.PullBlockStub = nil
	fake.pullBlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlockPuller) PullBlockReturnsOnCall(i int, result1 error) {
	fake.pullBlockMutex.Lock()
	defer fake.pullBlockMutex.Unlock()
	fake.PullBlockStub = nil
	if fake.pullBlockReturnsOnCall == nil {
		fake.pullBlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pullBlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlockPuller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pullBlockMutex.RLock()
	defer fake.pullBlockMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlockPuller) recordInvocation(key string, args []interface{}) {
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

var _ generate.BlockPuller = new(FakeBlockPuller)
