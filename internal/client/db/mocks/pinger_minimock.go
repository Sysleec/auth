package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/Sysleec/auth/internal/client/db.Pinger -o pinger_minimock.go -n PingerMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// PingerMock implements db.Pinger
type PingerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcPing          func(ctx context.Context) (err error)
	inspectFuncPing   func(ctx context.Context)
	afterPingCounter  uint64
	beforePingCounter uint64
	PingMock          mPingerMockPing
}

// NewPingerMock returns a mock for db.Pinger
func NewPingerMock(t minimock.Tester) *PingerMock {
	m := &PingerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.PingMock = mPingerMockPing{mock: m}
	m.PingMock.callArgs = []*PingerMockPingParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mPingerMockPing struct {
	mock               *PingerMock
	defaultExpectation *PingerMockPingExpectation
	expectations       []*PingerMockPingExpectation

	callArgs []*PingerMockPingParams
	mutex    sync.RWMutex
}

// PingerMockPingExpectation specifies expectation struct of the Pinger.Ping
type PingerMockPingExpectation struct {
	mock    *PingerMock
	params  *PingerMockPingParams
	results *PingerMockPingResults
	Counter uint64
}

// PingerMockPingParams contains parameters of the Pinger.Ping
type PingerMockPingParams struct {
	ctx context.Context
}

// PingerMockPingResults contains results of the Pinger.Ping
type PingerMockPingResults struct {
	err error
}

// Expect sets up expected params for Pinger.Ping
func (mmPing *mPingerMockPing) Expect(ctx context.Context) *mPingerMockPing {
	if mmPing.mock.funcPing != nil {
		mmPing.mock.t.Fatalf("PingerMock.Ping mock is already set by Set")
	}

	if mmPing.defaultExpectation == nil {
		mmPing.defaultExpectation = &PingerMockPingExpectation{}
	}

	mmPing.defaultExpectation.params = &PingerMockPingParams{ctx}
	for _, e := range mmPing.expectations {
		if minimock.Equal(e.params, mmPing.defaultExpectation.params) {
			mmPing.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPing.defaultExpectation.params)
		}
	}

	return mmPing
}

// Inspect accepts an inspector function that has same arguments as the Pinger.Ping
func (mmPing *mPingerMockPing) Inspect(f func(ctx context.Context)) *mPingerMockPing {
	if mmPing.mock.inspectFuncPing != nil {
		mmPing.mock.t.Fatalf("Inspect function is already set for PingerMock.Ping")
	}

	mmPing.mock.inspectFuncPing = f

	return mmPing
}

// Return sets up results that will be returned by Pinger.Ping
func (mmPing *mPingerMockPing) Return(err error) *PingerMock {
	if mmPing.mock.funcPing != nil {
		mmPing.mock.t.Fatalf("PingerMock.Ping mock is already set by Set")
	}

	if mmPing.defaultExpectation == nil {
		mmPing.defaultExpectation = &PingerMockPingExpectation{mock: mmPing.mock}
	}
	mmPing.defaultExpectation.results = &PingerMockPingResults{err}
	return mmPing.mock
}

// Set uses given function f to mock the Pinger.Ping method
func (mmPing *mPingerMockPing) Set(f func(ctx context.Context) (err error)) *PingerMock {
	if mmPing.defaultExpectation != nil {
		mmPing.mock.t.Fatalf("Default expectation is already set for the Pinger.Ping method")
	}

	if len(mmPing.expectations) > 0 {
		mmPing.mock.t.Fatalf("Some expectations are already set for the Pinger.Ping method")
	}

	mmPing.mock.funcPing = f
	return mmPing.mock
}

// When sets expectation for the Pinger.Ping which will trigger the result defined by the following
// Then helper
func (mmPing *mPingerMockPing) When(ctx context.Context) *PingerMockPingExpectation {
	if mmPing.mock.funcPing != nil {
		mmPing.mock.t.Fatalf("PingerMock.Ping mock is already set by Set")
	}

	expectation := &PingerMockPingExpectation{
		mock:   mmPing.mock,
		params: &PingerMockPingParams{ctx},
	}
	mmPing.expectations = append(mmPing.expectations, expectation)
	return expectation
}

// Then sets up Pinger.Ping return parameters for the expectation previously defined by the When method
func (e *PingerMockPingExpectation) Then(err error) *PingerMock {
	e.results = &PingerMockPingResults{err}
	return e.mock
}

// Ping implements db.Pinger
func (mmPing *PingerMock) Ping(ctx context.Context) (err error) {
	mm_atomic.AddUint64(&mmPing.beforePingCounter, 1)
	defer mm_atomic.AddUint64(&mmPing.afterPingCounter, 1)

	if mmPing.inspectFuncPing != nil {
		mmPing.inspectFuncPing(ctx)
	}

	mm_params := PingerMockPingParams{ctx}

	// Record call args
	mmPing.PingMock.mutex.Lock()
	mmPing.PingMock.callArgs = append(mmPing.PingMock.callArgs, &mm_params)
	mmPing.PingMock.mutex.Unlock()

	for _, e := range mmPing.PingMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmPing.PingMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPing.PingMock.defaultExpectation.Counter, 1)
		mm_want := mmPing.PingMock.defaultExpectation.params
		mm_got := PingerMockPingParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPing.t.Errorf("PingerMock.Ping got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPing.PingMock.defaultExpectation.results
		if mm_results == nil {
			mmPing.t.Fatal("No results are set for the PingerMock.Ping")
		}
		return (*mm_results).err
	}
	if mmPing.funcPing != nil {
		return mmPing.funcPing(ctx)
	}
	mmPing.t.Fatalf("Unexpected call to PingerMock.Ping. %v", ctx)
	return
}

// PingAfterCounter returns a count of finished PingerMock.Ping invocations
func (mmPing *PingerMock) PingAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPing.afterPingCounter)
}

// PingBeforeCounter returns a count of PingerMock.Ping invocations
func (mmPing *PingerMock) PingBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPing.beforePingCounter)
}

// Calls returns a list of arguments used in each call to PingerMock.Ping.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPing *mPingerMockPing) Calls() []*PingerMockPingParams {
	mmPing.mutex.RLock()

	argCopy := make([]*PingerMockPingParams, len(mmPing.callArgs))
	copy(argCopy, mmPing.callArgs)

	mmPing.mutex.RUnlock()

	return argCopy
}

// MinimockPingDone returns true if the count of the Ping invocations corresponds
// the number of defined expectations
func (m *PingerMock) MinimockPingDone() bool {
	for _, e := range m.PingMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PingMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPingCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPing != nil && mm_atomic.LoadUint64(&m.afterPingCounter) < 1 {
		return false
	}
	return true
}

// MinimockPingInspect logs each unmet expectation
func (m *PingerMock) MinimockPingInspect() {
	for _, e := range m.PingMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PingerMock.Ping with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PingMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPingCounter) < 1 {
		if m.PingMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PingerMock.Ping")
		} else {
			m.t.Errorf("Expected call to PingerMock.Ping with params: %#v", *m.PingMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPing != nil && mm_atomic.LoadUint64(&m.afterPingCounter) < 1 {
		m.t.Error("Expected call to PingerMock.Ping")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PingerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockPingInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PingerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *PingerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockPingDone()
}
