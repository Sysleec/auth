package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/Sysleec/auth/internal/client/db.QueryExecer -o query_execer_minimock.go -n QueryExecerMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	mm_db "github.com/Sysleec/auth/internal/client/db"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// QueryExecerMock implements db.QueryExecer
type QueryExecerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcExecContext          func(ctx context.Context, q mm_db.Query, args ...interface{}) (c2 pgconn.CommandTag, err error)
	inspectFuncExecContext   func(ctx context.Context, q mm_db.Query, args ...interface{})
	afterExecContextCounter  uint64
	beforeExecContextCounter uint64
	ExecContextMock          mQueryExecerMockExecContext

	funcQueryContext          func(ctx context.Context, q mm_db.Query, args ...interface{}) (r1 pgx.Rows, err error)
	inspectFuncQueryContext   func(ctx context.Context, q mm_db.Query, args ...interface{})
	afterQueryContextCounter  uint64
	beforeQueryContextCounter uint64
	QueryContextMock          mQueryExecerMockQueryContext

	funcQueryRowContext          func(ctx context.Context, q mm_db.Query, args ...interface{}) (r1 pgx.Row)
	inspectFuncQueryRowContext   func(ctx context.Context, q mm_db.Query, args ...interface{})
	afterQueryRowContextCounter  uint64
	beforeQueryRowContextCounter uint64
	QueryRowContextMock          mQueryExecerMockQueryRowContext
}

// NewQueryExecerMock returns a mock for db.QueryExecer
func NewQueryExecerMock(t minimock.Tester) *QueryExecerMock {
	m := &QueryExecerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ExecContextMock = mQueryExecerMockExecContext{mock: m}
	m.ExecContextMock.callArgs = []*QueryExecerMockExecContextParams{}

	m.QueryContextMock = mQueryExecerMockQueryContext{mock: m}
	m.QueryContextMock.callArgs = []*QueryExecerMockQueryContextParams{}

	m.QueryRowContextMock = mQueryExecerMockQueryRowContext{mock: m}
	m.QueryRowContextMock.callArgs = []*QueryExecerMockQueryRowContextParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mQueryExecerMockExecContext struct {
	mock               *QueryExecerMock
	defaultExpectation *QueryExecerMockExecContextExpectation
	expectations       []*QueryExecerMockExecContextExpectation

	callArgs []*QueryExecerMockExecContextParams
	mutex    sync.RWMutex
}

// QueryExecerMockExecContextExpectation specifies expectation struct of the QueryExecer.ExecContext
type QueryExecerMockExecContextExpectation struct {
	mock    *QueryExecerMock
	params  *QueryExecerMockExecContextParams
	results *QueryExecerMockExecContextResults
	Counter uint64
}

// QueryExecerMockExecContextParams contains parameters of the QueryExecer.ExecContext
type QueryExecerMockExecContextParams struct {
	ctx  context.Context
	q    mm_db.Query
	args []interface{}
}

// QueryExecerMockExecContextResults contains results of the QueryExecer.ExecContext
type QueryExecerMockExecContextResults struct {
	c2  pgconn.CommandTag
	err error
}

// Expect sets up expected params for QueryExecer.ExecContext
func (mmExecContext *mQueryExecerMockExecContext) Expect(ctx context.Context, q mm_db.Query, args ...interface{}) *mQueryExecerMockExecContext {
	if mmExecContext.mock.funcExecContext != nil {
		mmExecContext.mock.t.Fatalf("QueryExecerMock.ExecContext mock is already set by Set")
	}

	if mmExecContext.defaultExpectation == nil {
		mmExecContext.defaultExpectation = &QueryExecerMockExecContextExpectation{}
	}

	mmExecContext.defaultExpectation.params = &QueryExecerMockExecContextParams{ctx, q, args}
	for _, e := range mmExecContext.expectations {
		if minimock.Equal(e.params, mmExecContext.defaultExpectation.params) {
			mmExecContext.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmExecContext.defaultExpectation.params)
		}
	}

	return mmExecContext
}

// Inspect accepts an inspector function that has same arguments as the QueryExecer.ExecContext
func (mmExecContext *mQueryExecerMockExecContext) Inspect(f func(ctx context.Context, q mm_db.Query, args ...interface{})) *mQueryExecerMockExecContext {
	if mmExecContext.mock.inspectFuncExecContext != nil {
		mmExecContext.mock.t.Fatalf("Inspect function is already set for QueryExecerMock.ExecContext")
	}

	mmExecContext.mock.inspectFuncExecContext = f

	return mmExecContext
}

// Return sets up results that will be returned by QueryExecer.ExecContext
func (mmExecContext *mQueryExecerMockExecContext) Return(c2 pgconn.CommandTag, err error) *QueryExecerMock {
	if mmExecContext.mock.funcExecContext != nil {
		mmExecContext.mock.t.Fatalf("QueryExecerMock.ExecContext mock is already set by Set")
	}

	if mmExecContext.defaultExpectation == nil {
		mmExecContext.defaultExpectation = &QueryExecerMockExecContextExpectation{mock: mmExecContext.mock}
	}
	mmExecContext.defaultExpectation.results = &QueryExecerMockExecContextResults{c2, err}
	return mmExecContext.mock
}

// Set uses given function f to mock the QueryExecer.ExecContext method
func (mmExecContext *mQueryExecerMockExecContext) Set(f func(ctx context.Context, q mm_db.Query, args ...interface{}) (c2 pgconn.CommandTag, err error)) *QueryExecerMock {
	if mmExecContext.defaultExpectation != nil {
		mmExecContext.mock.t.Fatalf("Default expectation is already set for the QueryExecer.ExecContext method")
	}

	if len(mmExecContext.expectations) > 0 {
		mmExecContext.mock.t.Fatalf("Some expectations are already set for the QueryExecer.ExecContext method")
	}

	mmExecContext.mock.funcExecContext = f
	return mmExecContext.mock
}

// When sets expectation for the QueryExecer.ExecContext which will trigger the result defined by the following
// Then helper
func (mmExecContext *mQueryExecerMockExecContext) When(ctx context.Context, q mm_db.Query, args ...interface{}) *QueryExecerMockExecContextExpectation {
	if mmExecContext.mock.funcExecContext != nil {
		mmExecContext.mock.t.Fatalf("QueryExecerMock.ExecContext mock is already set by Set")
	}

	expectation := &QueryExecerMockExecContextExpectation{
		mock:   mmExecContext.mock,
		params: &QueryExecerMockExecContextParams{ctx, q, args},
	}
	mmExecContext.expectations = append(mmExecContext.expectations, expectation)
	return expectation
}

// Then sets up QueryExecer.ExecContext return parameters for the expectation previously defined by the When method
func (e *QueryExecerMockExecContextExpectation) Then(c2 pgconn.CommandTag, err error) *QueryExecerMock {
	e.results = &QueryExecerMockExecContextResults{c2, err}
	return e.mock
}

// ExecContext implements db.QueryExecer
func (mmExecContext *QueryExecerMock) ExecContext(ctx context.Context, q mm_db.Query, args ...interface{}) (c2 pgconn.CommandTag, err error) {
	mm_atomic.AddUint64(&mmExecContext.beforeExecContextCounter, 1)
	defer mm_atomic.AddUint64(&mmExecContext.afterExecContextCounter, 1)

	if mmExecContext.inspectFuncExecContext != nil {
		mmExecContext.inspectFuncExecContext(ctx, q, args...)
	}

	mm_params := QueryExecerMockExecContextParams{ctx, q, args}

	// Record call args
	mmExecContext.ExecContextMock.mutex.Lock()
	mmExecContext.ExecContextMock.callArgs = append(mmExecContext.ExecContextMock.callArgs, &mm_params)
	mmExecContext.ExecContextMock.mutex.Unlock()

	for _, e := range mmExecContext.ExecContextMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.c2, e.results.err
		}
	}

	if mmExecContext.ExecContextMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmExecContext.ExecContextMock.defaultExpectation.Counter, 1)
		mm_want := mmExecContext.ExecContextMock.defaultExpectation.params
		mm_got := QueryExecerMockExecContextParams{ctx, q, args}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmExecContext.t.Errorf("QueryExecerMock.ExecContext got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmExecContext.ExecContextMock.defaultExpectation.results
		if mm_results == nil {
			mmExecContext.t.Fatal("No results are set for the QueryExecerMock.ExecContext")
		}
		return (*mm_results).c2, (*mm_results).err
	}
	if mmExecContext.funcExecContext != nil {
		return mmExecContext.funcExecContext(ctx, q, args...)
	}
	mmExecContext.t.Fatalf("Unexpected call to QueryExecerMock.ExecContext. %v %v %v", ctx, q, args)
	return
}

// ExecContextAfterCounter returns a count of finished QueryExecerMock.ExecContext invocations
func (mmExecContext *QueryExecerMock) ExecContextAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExecContext.afterExecContextCounter)
}

// ExecContextBeforeCounter returns a count of QueryExecerMock.ExecContext invocations
func (mmExecContext *QueryExecerMock) ExecContextBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExecContext.beforeExecContextCounter)
}

// Calls returns a list of arguments used in each call to QueryExecerMock.ExecContext.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmExecContext *mQueryExecerMockExecContext) Calls() []*QueryExecerMockExecContextParams {
	mmExecContext.mutex.RLock()

	argCopy := make([]*QueryExecerMockExecContextParams, len(mmExecContext.callArgs))
	copy(argCopy, mmExecContext.callArgs)

	mmExecContext.mutex.RUnlock()

	return argCopy
}

// MinimockExecContextDone returns true if the count of the ExecContext invocations corresponds
// the number of defined expectations
func (m *QueryExecerMock) MinimockExecContextDone() bool {
	for _, e := range m.ExecContextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExecContextMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExecContextCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExecContext != nil && mm_atomic.LoadUint64(&m.afterExecContextCounter) < 1 {
		return false
	}
	return true
}

// MinimockExecContextInspect logs each unmet expectation
func (m *QueryExecerMock) MinimockExecContextInspect() {
	for _, e := range m.ExecContextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to QueryExecerMock.ExecContext with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExecContextMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExecContextCounter) < 1 {
		if m.ExecContextMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to QueryExecerMock.ExecContext")
		} else {
			m.t.Errorf("Expected call to QueryExecerMock.ExecContext with params: %#v", *m.ExecContextMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExecContext != nil && mm_atomic.LoadUint64(&m.afterExecContextCounter) < 1 {
		m.t.Error("Expected call to QueryExecerMock.ExecContext")
	}
}

type mQueryExecerMockQueryContext struct {
	mock               *QueryExecerMock
	defaultExpectation *QueryExecerMockQueryContextExpectation
	expectations       []*QueryExecerMockQueryContextExpectation

	callArgs []*QueryExecerMockQueryContextParams
	mutex    sync.RWMutex
}

// QueryExecerMockQueryContextExpectation specifies expectation struct of the QueryExecer.QueryContext
type QueryExecerMockQueryContextExpectation struct {
	mock    *QueryExecerMock
	params  *QueryExecerMockQueryContextParams
	results *QueryExecerMockQueryContextResults
	Counter uint64
}

// QueryExecerMockQueryContextParams contains parameters of the QueryExecer.QueryContext
type QueryExecerMockQueryContextParams struct {
	ctx  context.Context
	q    mm_db.Query
	args []interface{}
}

// QueryExecerMockQueryContextResults contains results of the QueryExecer.QueryContext
type QueryExecerMockQueryContextResults struct {
	r1  pgx.Rows
	err error
}

// Expect sets up expected params for QueryExecer.QueryContext
func (mmQueryContext *mQueryExecerMockQueryContext) Expect(ctx context.Context, q mm_db.Query, args ...interface{}) *mQueryExecerMockQueryContext {
	if mmQueryContext.mock.funcQueryContext != nil {
		mmQueryContext.mock.t.Fatalf("QueryExecerMock.QueryContext mock is already set by Set")
	}

	if mmQueryContext.defaultExpectation == nil {
		mmQueryContext.defaultExpectation = &QueryExecerMockQueryContextExpectation{}
	}

	mmQueryContext.defaultExpectation.params = &QueryExecerMockQueryContextParams{ctx, q, args}
	for _, e := range mmQueryContext.expectations {
		if minimock.Equal(e.params, mmQueryContext.defaultExpectation.params) {
			mmQueryContext.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmQueryContext.defaultExpectation.params)
		}
	}

	return mmQueryContext
}

// Inspect accepts an inspector function that has same arguments as the QueryExecer.QueryContext
func (mmQueryContext *mQueryExecerMockQueryContext) Inspect(f func(ctx context.Context, q mm_db.Query, args ...interface{})) *mQueryExecerMockQueryContext {
	if mmQueryContext.mock.inspectFuncQueryContext != nil {
		mmQueryContext.mock.t.Fatalf("Inspect function is already set for QueryExecerMock.QueryContext")
	}

	mmQueryContext.mock.inspectFuncQueryContext = f

	return mmQueryContext
}

// Return sets up results that will be returned by QueryExecer.QueryContext
func (mmQueryContext *mQueryExecerMockQueryContext) Return(r1 pgx.Rows, err error) *QueryExecerMock {
	if mmQueryContext.mock.funcQueryContext != nil {
		mmQueryContext.mock.t.Fatalf("QueryExecerMock.QueryContext mock is already set by Set")
	}

	if mmQueryContext.defaultExpectation == nil {
		mmQueryContext.defaultExpectation = &QueryExecerMockQueryContextExpectation{mock: mmQueryContext.mock}
	}
	mmQueryContext.defaultExpectation.results = &QueryExecerMockQueryContextResults{r1, err}
	return mmQueryContext.mock
}

// Set uses given function f to mock the QueryExecer.QueryContext method
func (mmQueryContext *mQueryExecerMockQueryContext) Set(f func(ctx context.Context, q mm_db.Query, args ...interface{}) (r1 pgx.Rows, err error)) *QueryExecerMock {
	if mmQueryContext.defaultExpectation != nil {
		mmQueryContext.mock.t.Fatalf("Default expectation is already set for the QueryExecer.QueryContext method")
	}

	if len(mmQueryContext.expectations) > 0 {
		mmQueryContext.mock.t.Fatalf("Some expectations are already set for the QueryExecer.QueryContext method")
	}

	mmQueryContext.mock.funcQueryContext = f
	return mmQueryContext.mock
}

// When sets expectation for the QueryExecer.QueryContext which will trigger the result defined by the following
// Then helper
func (mmQueryContext *mQueryExecerMockQueryContext) When(ctx context.Context, q mm_db.Query, args ...interface{}) *QueryExecerMockQueryContextExpectation {
	if mmQueryContext.mock.funcQueryContext != nil {
		mmQueryContext.mock.t.Fatalf("QueryExecerMock.QueryContext mock is already set by Set")
	}

	expectation := &QueryExecerMockQueryContextExpectation{
		mock:   mmQueryContext.mock,
		params: &QueryExecerMockQueryContextParams{ctx, q, args},
	}
	mmQueryContext.expectations = append(mmQueryContext.expectations, expectation)
	return expectation
}

// Then sets up QueryExecer.QueryContext return parameters for the expectation previously defined by the When method
func (e *QueryExecerMockQueryContextExpectation) Then(r1 pgx.Rows, err error) *QueryExecerMock {
	e.results = &QueryExecerMockQueryContextResults{r1, err}
	return e.mock
}

// QueryContext implements db.QueryExecer
func (mmQueryContext *QueryExecerMock) QueryContext(ctx context.Context, q mm_db.Query, args ...interface{}) (r1 pgx.Rows, err error) {
	mm_atomic.AddUint64(&mmQueryContext.beforeQueryContextCounter, 1)
	defer mm_atomic.AddUint64(&mmQueryContext.afterQueryContextCounter, 1)

	if mmQueryContext.inspectFuncQueryContext != nil {
		mmQueryContext.inspectFuncQueryContext(ctx, q, args...)
	}

	mm_params := QueryExecerMockQueryContextParams{ctx, q, args}

	// Record call args
	mmQueryContext.QueryContextMock.mutex.Lock()
	mmQueryContext.QueryContextMock.callArgs = append(mmQueryContext.QueryContextMock.callArgs, &mm_params)
	mmQueryContext.QueryContextMock.mutex.Unlock()

	for _, e := range mmQueryContext.QueryContextMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1, e.results.err
		}
	}

	if mmQueryContext.QueryContextMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmQueryContext.QueryContextMock.defaultExpectation.Counter, 1)
		mm_want := mmQueryContext.QueryContextMock.defaultExpectation.params
		mm_got := QueryExecerMockQueryContextParams{ctx, q, args}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmQueryContext.t.Errorf("QueryExecerMock.QueryContext got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmQueryContext.QueryContextMock.defaultExpectation.results
		if mm_results == nil {
			mmQueryContext.t.Fatal("No results are set for the QueryExecerMock.QueryContext")
		}
		return (*mm_results).r1, (*mm_results).err
	}
	if mmQueryContext.funcQueryContext != nil {
		return mmQueryContext.funcQueryContext(ctx, q, args...)
	}
	mmQueryContext.t.Fatalf("Unexpected call to QueryExecerMock.QueryContext. %v %v %v", ctx, q, args)
	return
}

// QueryContextAfterCounter returns a count of finished QueryExecerMock.QueryContext invocations
func (mmQueryContext *QueryExecerMock) QueryContextAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmQueryContext.afterQueryContextCounter)
}

// QueryContextBeforeCounter returns a count of QueryExecerMock.QueryContext invocations
func (mmQueryContext *QueryExecerMock) QueryContextBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmQueryContext.beforeQueryContextCounter)
}

// Calls returns a list of arguments used in each call to QueryExecerMock.QueryContext.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmQueryContext *mQueryExecerMockQueryContext) Calls() []*QueryExecerMockQueryContextParams {
	mmQueryContext.mutex.RLock()

	argCopy := make([]*QueryExecerMockQueryContextParams, len(mmQueryContext.callArgs))
	copy(argCopy, mmQueryContext.callArgs)

	mmQueryContext.mutex.RUnlock()

	return argCopy
}

// MinimockQueryContextDone returns true if the count of the QueryContext invocations corresponds
// the number of defined expectations
func (m *QueryExecerMock) MinimockQueryContextDone() bool {
	for _, e := range m.QueryContextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.QueryContextMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterQueryContextCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcQueryContext != nil && mm_atomic.LoadUint64(&m.afterQueryContextCounter) < 1 {
		return false
	}
	return true
}

// MinimockQueryContextInspect logs each unmet expectation
func (m *QueryExecerMock) MinimockQueryContextInspect() {
	for _, e := range m.QueryContextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to QueryExecerMock.QueryContext with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.QueryContextMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterQueryContextCounter) < 1 {
		if m.QueryContextMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to QueryExecerMock.QueryContext")
		} else {
			m.t.Errorf("Expected call to QueryExecerMock.QueryContext with params: %#v", *m.QueryContextMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcQueryContext != nil && mm_atomic.LoadUint64(&m.afterQueryContextCounter) < 1 {
		m.t.Error("Expected call to QueryExecerMock.QueryContext")
	}
}

type mQueryExecerMockQueryRowContext struct {
	mock               *QueryExecerMock
	defaultExpectation *QueryExecerMockQueryRowContextExpectation
	expectations       []*QueryExecerMockQueryRowContextExpectation

	callArgs []*QueryExecerMockQueryRowContextParams
	mutex    sync.RWMutex
}

// QueryExecerMockQueryRowContextExpectation specifies expectation struct of the QueryExecer.QueryRowContext
type QueryExecerMockQueryRowContextExpectation struct {
	mock    *QueryExecerMock
	params  *QueryExecerMockQueryRowContextParams
	results *QueryExecerMockQueryRowContextResults
	Counter uint64
}

// QueryExecerMockQueryRowContextParams contains parameters of the QueryExecer.QueryRowContext
type QueryExecerMockQueryRowContextParams struct {
	ctx  context.Context
	q    mm_db.Query
	args []interface{}
}

// QueryExecerMockQueryRowContextResults contains results of the QueryExecer.QueryRowContext
type QueryExecerMockQueryRowContextResults struct {
	r1 pgx.Row
}

// Expect sets up expected params for QueryExecer.QueryRowContext
func (mmQueryRowContext *mQueryExecerMockQueryRowContext) Expect(ctx context.Context, q mm_db.Query, args ...interface{}) *mQueryExecerMockQueryRowContext {
	if mmQueryRowContext.mock.funcQueryRowContext != nil {
		mmQueryRowContext.mock.t.Fatalf("QueryExecerMock.QueryRowContext mock is already set by Set")
	}

	if mmQueryRowContext.defaultExpectation == nil {
		mmQueryRowContext.defaultExpectation = &QueryExecerMockQueryRowContextExpectation{}
	}

	mmQueryRowContext.defaultExpectation.params = &QueryExecerMockQueryRowContextParams{ctx, q, args}
	for _, e := range mmQueryRowContext.expectations {
		if minimock.Equal(e.params, mmQueryRowContext.defaultExpectation.params) {
			mmQueryRowContext.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmQueryRowContext.defaultExpectation.params)
		}
	}

	return mmQueryRowContext
}

// Inspect accepts an inspector function that has same arguments as the QueryExecer.QueryRowContext
func (mmQueryRowContext *mQueryExecerMockQueryRowContext) Inspect(f func(ctx context.Context, q mm_db.Query, args ...interface{})) *mQueryExecerMockQueryRowContext {
	if mmQueryRowContext.mock.inspectFuncQueryRowContext != nil {
		mmQueryRowContext.mock.t.Fatalf("Inspect function is already set for QueryExecerMock.QueryRowContext")
	}

	mmQueryRowContext.mock.inspectFuncQueryRowContext = f

	return mmQueryRowContext
}

// Return sets up results that will be returned by QueryExecer.QueryRowContext
func (mmQueryRowContext *mQueryExecerMockQueryRowContext) Return(r1 pgx.Row) *QueryExecerMock {
	if mmQueryRowContext.mock.funcQueryRowContext != nil {
		mmQueryRowContext.mock.t.Fatalf("QueryExecerMock.QueryRowContext mock is already set by Set")
	}

	if mmQueryRowContext.defaultExpectation == nil {
		mmQueryRowContext.defaultExpectation = &QueryExecerMockQueryRowContextExpectation{mock: mmQueryRowContext.mock}
	}
	mmQueryRowContext.defaultExpectation.results = &QueryExecerMockQueryRowContextResults{r1}
	return mmQueryRowContext.mock
}

// Set uses given function f to mock the QueryExecer.QueryRowContext method
func (mmQueryRowContext *mQueryExecerMockQueryRowContext) Set(f func(ctx context.Context, q mm_db.Query, args ...interface{}) (r1 pgx.Row)) *QueryExecerMock {
	if mmQueryRowContext.defaultExpectation != nil {
		mmQueryRowContext.mock.t.Fatalf("Default expectation is already set for the QueryExecer.QueryRowContext method")
	}

	if len(mmQueryRowContext.expectations) > 0 {
		mmQueryRowContext.mock.t.Fatalf("Some expectations are already set for the QueryExecer.QueryRowContext method")
	}

	mmQueryRowContext.mock.funcQueryRowContext = f
	return mmQueryRowContext.mock
}

// When sets expectation for the QueryExecer.QueryRowContext which will trigger the result defined by the following
// Then helper
func (mmQueryRowContext *mQueryExecerMockQueryRowContext) When(ctx context.Context, q mm_db.Query, args ...interface{}) *QueryExecerMockQueryRowContextExpectation {
	if mmQueryRowContext.mock.funcQueryRowContext != nil {
		mmQueryRowContext.mock.t.Fatalf("QueryExecerMock.QueryRowContext mock is already set by Set")
	}

	expectation := &QueryExecerMockQueryRowContextExpectation{
		mock:   mmQueryRowContext.mock,
		params: &QueryExecerMockQueryRowContextParams{ctx, q, args},
	}
	mmQueryRowContext.expectations = append(mmQueryRowContext.expectations, expectation)
	return expectation
}

// Then sets up QueryExecer.QueryRowContext return parameters for the expectation previously defined by the When method
func (e *QueryExecerMockQueryRowContextExpectation) Then(r1 pgx.Row) *QueryExecerMock {
	e.results = &QueryExecerMockQueryRowContextResults{r1}
	return e.mock
}

// QueryRowContext implements db.QueryExecer
func (mmQueryRowContext *QueryExecerMock) QueryRowContext(ctx context.Context, q mm_db.Query, args ...interface{}) (r1 pgx.Row) {
	mm_atomic.AddUint64(&mmQueryRowContext.beforeQueryRowContextCounter, 1)
	defer mm_atomic.AddUint64(&mmQueryRowContext.afterQueryRowContextCounter, 1)

	if mmQueryRowContext.inspectFuncQueryRowContext != nil {
		mmQueryRowContext.inspectFuncQueryRowContext(ctx, q, args...)
	}

	mm_params := QueryExecerMockQueryRowContextParams{ctx, q, args}

	// Record call args
	mmQueryRowContext.QueryRowContextMock.mutex.Lock()
	mmQueryRowContext.QueryRowContextMock.callArgs = append(mmQueryRowContext.QueryRowContextMock.callArgs, &mm_params)
	mmQueryRowContext.QueryRowContextMock.mutex.Unlock()

	for _, e := range mmQueryRowContext.QueryRowContextMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1
		}
	}

	if mmQueryRowContext.QueryRowContextMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmQueryRowContext.QueryRowContextMock.defaultExpectation.Counter, 1)
		mm_want := mmQueryRowContext.QueryRowContextMock.defaultExpectation.params
		mm_got := QueryExecerMockQueryRowContextParams{ctx, q, args}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmQueryRowContext.t.Errorf("QueryExecerMock.QueryRowContext got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmQueryRowContext.QueryRowContextMock.defaultExpectation.results
		if mm_results == nil {
			mmQueryRowContext.t.Fatal("No results are set for the QueryExecerMock.QueryRowContext")
		}
		return (*mm_results).r1
	}
	if mmQueryRowContext.funcQueryRowContext != nil {
		return mmQueryRowContext.funcQueryRowContext(ctx, q, args...)
	}
	mmQueryRowContext.t.Fatalf("Unexpected call to QueryExecerMock.QueryRowContext. %v %v %v", ctx, q, args)
	return
}

// QueryRowContextAfterCounter returns a count of finished QueryExecerMock.QueryRowContext invocations
func (mmQueryRowContext *QueryExecerMock) QueryRowContextAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmQueryRowContext.afterQueryRowContextCounter)
}

// QueryRowContextBeforeCounter returns a count of QueryExecerMock.QueryRowContext invocations
func (mmQueryRowContext *QueryExecerMock) QueryRowContextBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmQueryRowContext.beforeQueryRowContextCounter)
}

// Calls returns a list of arguments used in each call to QueryExecerMock.QueryRowContext.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmQueryRowContext *mQueryExecerMockQueryRowContext) Calls() []*QueryExecerMockQueryRowContextParams {
	mmQueryRowContext.mutex.RLock()

	argCopy := make([]*QueryExecerMockQueryRowContextParams, len(mmQueryRowContext.callArgs))
	copy(argCopy, mmQueryRowContext.callArgs)

	mmQueryRowContext.mutex.RUnlock()

	return argCopy
}

// MinimockQueryRowContextDone returns true if the count of the QueryRowContext invocations corresponds
// the number of defined expectations
func (m *QueryExecerMock) MinimockQueryRowContextDone() bool {
	for _, e := range m.QueryRowContextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.QueryRowContextMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterQueryRowContextCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcQueryRowContext != nil && mm_atomic.LoadUint64(&m.afterQueryRowContextCounter) < 1 {
		return false
	}
	return true
}

// MinimockQueryRowContextInspect logs each unmet expectation
func (m *QueryExecerMock) MinimockQueryRowContextInspect() {
	for _, e := range m.QueryRowContextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to QueryExecerMock.QueryRowContext with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.QueryRowContextMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterQueryRowContextCounter) < 1 {
		if m.QueryRowContextMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to QueryExecerMock.QueryRowContext")
		} else {
			m.t.Errorf("Expected call to QueryExecerMock.QueryRowContext with params: %#v", *m.QueryRowContextMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcQueryRowContext != nil && mm_atomic.LoadUint64(&m.afterQueryRowContextCounter) < 1 {
		m.t.Error("Expected call to QueryExecerMock.QueryRowContext")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *QueryExecerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockExecContextInspect()

			m.MinimockQueryContextInspect()

			m.MinimockQueryRowContextInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *QueryExecerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *QueryExecerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockExecContextDone() &&
		m.MinimockQueryContextDone() &&
		m.MinimockQueryRowContextDone()
}
