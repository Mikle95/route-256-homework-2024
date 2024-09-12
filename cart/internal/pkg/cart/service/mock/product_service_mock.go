// Code generated by http://github.com/gojuno/minimock (v3.4.0). DO NOT EDIT.

package mock

//go:generate minimock -i gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/service.ProductService -o product_service_mock_test.go -n ProductServiceMock -p mock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

// ProductServiceMock implements mm_service.ProductService
type ProductServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetProduct          func(ctx context.Context, sku model.Sku) (ip1 *domain.Item, err error)
	funcGetProductOrigin    string
	inspectFuncGetProduct   func(ctx context.Context, sku model.Sku)
	afterGetProductCounter  uint64
	beforeGetProductCounter uint64
	GetProductMock          mProductServiceMockGetProduct
}

// NewProductServiceMock returns a mock for mm_service.ProductService
func NewProductServiceMock(t minimock.Tester) *ProductServiceMock {
	m := &ProductServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetProductMock = mProductServiceMockGetProduct{mock: m}
	m.GetProductMock.callArgs = []*ProductServiceMockGetProductParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mProductServiceMockGetProduct struct {
	optional           bool
	mock               *ProductServiceMock
	defaultExpectation *ProductServiceMockGetProductExpectation
	expectations       []*ProductServiceMockGetProductExpectation

	callArgs []*ProductServiceMockGetProductParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// ProductServiceMockGetProductExpectation specifies expectation struct of the ProductService.GetProduct
type ProductServiceMockGetProductExpectation struct {
	mock               *ProductServiceMock
	params             *ProductServiceMockGetProductParams
	paramPtrs          *ProductServiceMockGetProductParamPtrs
	expectationOrigins ProductServiceMockGetProductExpectationOrigins
	results            *ProductServiceMockGetProductResults
	returnOrigin       string
	Counter            uint64
}

// ProductServiceMockGetProductParams contains parameters of the ProductService.GetProduct
type ProductServiceMockGetProductParams struct {
	ctx context.Context
	sku model.Sku
}

// ProductServiceMockGetProductParamPtrs contains pointers to parameters of the ProductService.GetProduct
type ProductServiceMockGetProductParamPtrs struct {
	ctx *context.Context
	sku *model.Sku
}

// ProductServiceMockGetProductResults contains results of the ProductService.GetProduct
type ProductServiceMockGetProductResults struct {
	ip1 *domain.Item
	err error
}

// ProductServiceMockGetProductOrigins contains origins of expectations of the ProductService.GetProduct
type ProductServiceMockGetProductExpectationOrigins struct {
	origin    string
	originCtx string
	originSku string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGetProduct *mProductServiceMockGetProduct) Optional() *mProductServiceMockGetProduct {
	mmGetProduct.optional = true
	return mmGetProduct
}

// Expect sets up expected params for ProductService.GetProduct
func (mmGetProduct *mProductServiceMockGetProduct) Expect(ctx context.Context, sku model.Sku) *mProductServiceMockGetProduct {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceMock.GetProduct mock is already set by Set")
	}

	if mmGetProduct.defaultExpectation == nil {
		mmGetProduct.defaultExpectation = &ProductServiceMockGetProductExpectation{}
	}

	if mmGetProduct.defaultExpectation.paramPtrs != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceMock.GetProduct mock is already set by ExpectParams functions")
	}

	mmGetProduct.defaultExpectation.params = &ProductServiceMockGetProductParams{ctx, sku}
	mmGetProduct.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmGetProduct.expectations {
		if minimock.Equal(e.params, mmGetProduct.defaultExpectation.params) {
			mmGetProduct.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetProduct.defaultExpectation.params)
		}
	}

	return mmGetProduct
}

// ExpectCtxParam1 sets up expected param ctx for ProductService.GetProduct
func (mmGetProduct *mProductServiceMockGetProduct) ExpectCtxParam1(ctx context.Context) *mProductServiceMockGetProduct {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceMock.GetProduct mock is already set by Set")
	}

	if mmGetProduct.defaultExpectation == nil {
		mmGetProduct.defaultExpectation = &ProductServiceMockGetProductExpectation{}
	}

	if mmGetProduct.defaultExpectation.params != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceMock.GetProduct mock is already set by Expect")
	}

	if mmGetProduct.defaultExpectation.paramPtrs == nil {
		mmGetProduct.defaultExpectation.paramPtrs = &ProductServiceMockGetProductParamPtrs{}
	}
	mmGetProduct.defaultExpectation.paramPtrs.ctx = &ctx
	mmGetProduct.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmGetProduct
}

// ExpectSkuParam2 sets up expected param sku for ProductService.GetProduct
func (mmGetProduct *mProductServiceMockGetProduct) ExpectSkuParam2(sku model.Sku) *mProductServiceMockGetProduct {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceMock.GetProduct mock is already set by Set")
	}

	if mmGetProduct.defaultExpectation == nil {
		mmGetProduct.defaultExpectation = &ProductServiceMockGetProductExpectation{}
	}

	if mmGetProduct.defaultExpectation.params != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceMock.GetProduct mock is already set by Expect")
	}

	if mmGetProduct.defaultExpectation.paramPtrs == nil {
		mmGetProduct.defaultExpectation.paramPtrs = &ProductServiceMockGetProductParamPtrs{}
	}
	mmGetProduct.defaultExpectation.paramPtrs.sku = &sku
	mmGetProduct.defaultExpectation.expectationOrigins.originSku = minimock.CallerInfo(1)

	return mmGetProduct
}

// Inspect accepts an inspector function that has same arguments as the ProductService.GetProduct
func (mmGetProduct *mProductServiceMockGetProduct) Inspect(f func(ctx context.Context, sku model.Sku)) *mProductServiceMockGetProduct {
	if mmGetProduct.mock.inspectFuncGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("Inspect function is already set for ProductServiceMock.GetProduct")
	}

	mmGetProduct.mock.inspectFuncGetProduct = f

	return mmGetProduct
}

// Return sets up results that will be returned by ProductService.GetProduct
func (mmGetProduct *mProductServiceMockGetProduct) Return(ip1 *domain.Item, err error) *ProductServiceMock {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceMock.GetProduct mock is already set by Set")
	}

	if mmGetProduct.defaultExpectation == nil {
		mmGetProduct.defaultExpectation = &ProductServiceMockGetProductExpectation{mock: mmGetProduct.mock}
	}
	mmGetProduct.defaultExpectation.results = &ProductServiceMockGetProductResults{ip1, err}
	mmGetProduct.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGetProduct.mock
}

// Set uses given function f to mock the ProductService.GetProduct method
func (mmGetProduct *mProductServiceMockGetProduct) Set(f func(ctx context.Context, sku model.Sku) (ip1 *domain.Item, err error)) *ProductServiceMock {
	if mmGetProduct.defaultExpectation != nil {
		mmGetProduct.mock.t.Fatalf("Default expectation is already set for the ProductService.GetProduct method")
	}

	if len(mmGetProduct.expectations) > 0 {
		mmGetProduct.mock.t.Fatalf("Some expectations are already set for the ProductService.GetProduct method")
	}

	mmGetProduct.mock.funcGetProduct = f
	mmGetProduct.mock.funcGetProductOrigin = minimock.CallerInfo(1)
	return mmGetProduct.mock
}

// When sets expectation for the ProductService.GetProduct which will trigger the result defined by the following
// Then helper
func (mmGetProduct *mProductServiceMockGetProduct) When(ctx context.Context, sku model.Sku) *ProductServiceMockGetProductExpectation {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductServiceMock.GetProduct mock is already set by Set")
	}

	expectation := &ProductServiceMockGetProductExpectation{
		mock:               mmGetProduct.mock,
		params:             &ProductServiceMockGetProductParams{ctx, sku},
		expectationOrigins: ProductServiceMockGetProductExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmGetProduct.expectations = append(mmGetProduct.expectations, expectation)
	return expectation
}

// Then sets up ProductService.GetProduct return parameters for the expectation previously defined by the When method
func (e *ProductServiceMockGetProductExpectation) Then(ip1 *domain.Item, err error) *ProductServiceMock {
	e.results = &ProductServiceMockGetProductResults{ip1, err}
	return e.mock
}

// Times sets number of times ProductService.GetProduct should be invoked
func (mmGetProduct *mProductServiceMockGetProduct) Times(n uint64) *mProductServiceMockGetProduct {
	if n == 0 {
		mmGetProduct.mock.t.Fatalf("Times of ProductServiceMock.GetProduct mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGetProduct.expectedInvocations, n)
	mmGetProduct.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGetProduct
}

func (mmGetProduct *mProductServiceMockGetProduct) invocationsDone() bool {
	if len(mmGetProduct.expectations) == 0 && mmGetProduct.defaultExpectation == nil && mmGetProduct.mock.funcGetProduct == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGetProduct.mock.afterGetProductCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGetProduct.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// GetProduct implements mm_service.ProductService
func (mmGetProduct *ProductServiceMock) GetProduct(ctx context.Context, sku model.Sku) (ip1 *domain.Item, err error) {
	mm_atomic.AddUint64(&mmGetProduct.beforeGetProductCounter, 1)
	defer mm_atomic.AddUint64(&mmGetProduct.afterGetProductCounter, 1)

	mmGetProduct.t.Helper()

	if mmGetProduct.inspectFuncGetProduct != nil {
		mmGetProduct.inspectFuncGetProduct(ctx, sku)
	}

	mm_params := ProductServiceMockGetProductParams{ctx, sku}

	// Record call args
	mmGetProduct.GetProductMock.mutex.Lock()
	mmGetProduct.GetProductMock.callArgs = append(mmGetProduct.GetProductMock.callArgs, &mm_params)
	mmGetProduct.GetProductMock.mutex.Unlock()

	for _, e := range mmGetProduct.GetProductMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ip1, e.results.err
		}
	}

	if mmGetProduct.GetProductMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetProduct.GetProductMock.defaultExpectation.Counter, 1)
		mm_want := mmGetProduct.GetProductMock.defaultExpectation.params
		mm_want_ptrs := mmGetProduct.GetProductMock.defaultExpectation.paramPtrs

		mm_got := ProductServiceMockGetProductParams{ctx, sku}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGetProduct.t.Errorf("ProductServiceMock.GetProduct got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetProduct.GetProductMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.sku != nil && !minimock.Equal(*mm_want_ptrs.sku, mm_got.sku) {
				mmGetProduct.t.Errorf("ProductServiceMock.GetProduct got unexpected parameter sku, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetProduct.GetProductMock.defaultExpectation.expectationOrigins.originSku, *mm_want_ptrs.sku, mm_got.sku, minimock.Diff(*mm_want_ptrs.sku, mm_got.sku))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetProduct.t.Errorf("ProductServiceMock.GetProduct got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmGetProduct.GetProductMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetProduct.GetProductMock.defaultExpectation.results
		if mm_results == nil {
			mmGetProduct.t.Fatal("No results are set for the ProductServiceMock.GetProduct")
		}
		return (*mm_results).ip1, (*mm_results).err
	}
	if mmGetProduct.funcGetProduct != nil {
		return mmGetProduct.funcGetProduct(ctx, sku)
	}
	mmGetProduct.t.Fatalf("Unexpected call to ProductServiceMock.GetProduct. %v %v", ctx, sku)
	return
}

// GetProductAfterCounter returns a count of finished ProductServiceMock.GetProduct invocations
func (mmGetProduct *ProductServiceMock) GetProductAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetProduct.afterGetProductCounter)
}

// GetProductBeforeCounter returns a count of ProductServiceMock.GetProduct invocations
func (mmGetProduct *ProductServiceMock) GetProductBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetProduct.beforeGetProductCounter)
}

// Calls returns a list of arguments used in each call to ProductServiceMock.GetProduct.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetProduct *mProductServiceMockGetProduct) Calls() []*ProductServiceMockGetProductParams {
	mmGetProduct.mutex.RLock()

	argCopy := make([]*ProductServiceMockGetProductParams, len(mmGetProduct.callArgs))
	copy(argCopy, mmGetProduct.callArgs)

	mmGetProduct.mutex.RUnlock()

	return argCopy
}

// MinimockGetProductDone returns true if the count of the GetProduct invocations corresponds
// the number of defined expectations
func (m *ProductServiceMock) MinimockGetProductDone() bool {
	if m.GetProductMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetProductMock.invocationsDone()
}

// MinimockGetProductInspect logs each unmet expectation
func (m *ProductServiceMock) MinimockGetProductInspect() {
	for _, e := range m.GetProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ProductServiceMock.GetProduct at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterGetProductCounter := mm_atomic.LoadUint64(&m.afterGetProductCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetProductMock.defaultExpectation != nil && afterGetProductCounter < 1 {
		if m.GetProductMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to ProductServiceMock.GetProduct at\n%s", m.GetProductMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to ProductServiceMock.GetProduct at\n%s with params: %#v", m.GetProductMock.defaultExpectation.expectationOrigins.origin, *m.GetProductMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetProduct != nil && afterGetProductCounter < 1 {
		m.t.Errorf("Expected call to ProductServiceMock.GetProduct at\n%s", m.funcGetProductOrigin)
	}

	if !m.GetProductMock.invocationsDone() && afterGetProductCounter > 0 {
		m.t.Errorf("Expected %d calls to ProductServiceMock.GetProduct at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GetProductMock.expectedInvocations), m.GetProductMock.expectedInvocationsOrigin, afterGetProductCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ProductServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetProductInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ProductServiceMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ProductServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetProductDone()
}
