// Code generated by counterfeiter. DO NOT EDIT.
package courierfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/courier"
	"github.com/concourse/concourse/atc/db/lock"
)

type FakeMigrator struct {
	AcquireMigrationLockStub        func(lager.Logger) (lock.Lock, bool, error)
	acquireMigrationLockMutex       sync.RWMutex
	acquireMigrationLockArgsForCall []struct {
		arg1 lager.Logger
	}
	acquireMigrationLockReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	acquireMigrationLockReturnsOnCall map[int]struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	MigrateStub        func(lager.Logger) error
	migrateMutex       sync.RWMutex
	migrateArgsForCall []struct {
		arg1 lager.Logger
	}
	migrateReturns struct {
		result1 error
	}
	migrateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMigrator) AcquireMigrationLock(arg1 lager.Logger) (lock.Lock, bool, error) {
	fake.acquireMigrationLockMutex.Lock()
	ret, specificReturn := fake.acquireMigrationLockReturnsOnCall[len(fake.acquireMigrationLockArgsForCall)]
	fake.acquireMigrationLockArgsForCall = append(fake.acquireMigrationLockArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("AcquireMigrationLock", []interface{}{arg1})
	fake.acquireMigrationLockMutex.Unlock()
	if fake.AcquireMigrationLockStub != nil {
		return fake.AcquireMigrationLockStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.acquireMigrationLockReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeMigrator) AcquireMigrationLockCallCount() int {
	fake.acquireMigrationLockMutex.RLock()
	defer fake.acquireMigrationLockMutex.RUnlock()
	return len(fake.acquireMigrationLockArgsForCall)
}

func (fake *FakeMigrator) AcquireMigrationLockCalls(stub func(lager.Logger) (lock.Lock, bool, error)) {
	fake.acquireMigrationLockMutex.Lock()
	defer fake.acquireMigrationLockMutex.Unlock()
	fake.AcquireMigrationLockStub = stub
}

func (fake *FakeMigrator) AcquireMigrationLockArgsForCall(i int) lager.Logger {
	fake.acquireMigrationLockMutex.RLock()
	defer fake.acquireMigrationLockMutex.RUnlock()
	argsForCall := fake.acquireMigrationLockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMigrator) AcquireMigrationLockReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.acquireMigrationLockMutex.Lock()
	defer fake.acquireMigrationLockMutex.Unlock()
	fake.AcquireMigrationLockStub = nil
	fake.acquireMigrationLockReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeMigrator) AcquireMigrationLockReturnsOnCall(i int, result1 lock.Lock, result2 bool, result3 error) {
	fake.acquireMigrationLockMutex.Lock()
	defer fake.acquireMigrationLockMutex.Unlock()
	fake.AcquireMigrationLockStub = nil
	if fake.acquireMigrationLockReturnsOnCall == nil {
		fake.acquireMigrationLockReturnsOnCall = make(map[int]struct {
			result1 lock.Lock
			result2 bool
			result3 error
		})
	}
	fake.acquireMigrationLockReturnsOnCall[i] = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeMigrator) Migrate(arg1 lager.Logger) error {
	fake.migrateMutex.Lock()
	ret, specificReturn := fake.migrateReturnsOnCall[len(fake.migrateArgsForCall)]
	fake.migrateArgsForCall = append(fake.migrateArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Migrate", []interface{}{arg1})
	fake.migrateMutex.Unlock()
	if fake.MigrateStub != nil {
		return fake.MigrateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.migrateReturns
	return fakeReturns.result1
}

func (fake *FakeMigrator) MigrateCallCount() int {
	fake.migrateMutex.RLock()
	defer fake.migrateMutex.RUnlock()
	return len(fake.migrateArgsForCall)
}

func (fake *FakeMigrator) MigrateCalls(stub func(lager.Logger) error) {
	fake.migrateMutex.Lock()
	defer fake.migrateMutex.Unlock()
	fake.MigrateStub = stub
}

func (fake *FakeMigrator) MigrateArgsForCall(i int) lager.Logger {
	fake.migrateMutex.RLock()
	defer fake.migrateMutex.RUnlock()
	argsForCall := fake.migrateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMigrator) MigrateReturns(result1 error) {
	fake.migrateMutex.Lock()
	defer fake.migrateMutex.Unlock()
	fake.MigrateStub = nil
	fake.migrateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMigrator) MigrateReturnsOnCall(i int, result1 error) {
	fake.migrateMutex.Lock()
	defer fake.migrateMutex.Unlock()
	fake.MigrateStub = nil
	if fake.migrateReturnsOnCall == nil {
		fake.migrateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.migrateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMigrator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acquireMigrationLockMutex.RLock()
	defer fake.acquireMigrationLockMutex.RUnlock()
	fake.migrateMutex.RLock()
	defer fake.migrateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMigrator) recordInvocation(key string, args []interface{}) {
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

var _ courier.Migrator = new(FakeMigrator)