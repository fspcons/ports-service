// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package port

import (
	"context"
	"github.com/fspcons/ports-service/src/domain"
	"sync"
)

// Ensure, that UseCaseMock does implement UseCase.
// If this is not the case, regenerate this file with moq.
var _ UseCase = &UseCaseMock{}

// UseCaseMock is a mock implementation of UseCase.
//
//	func TestSomethingThatUsesUseCase(t *testing.T) {
//
//		// make and configure a mocked UseCase
//		mockedUseCase := &UseCaseMock{
//			CreateFunc: func(ctx context.Context, p *domain.Port) error {
//				panic("mock out the Create method")
//			},
//			UpdateFunc: func(ctx context.Context, id string, upd Update) (*domain.Port, error) {
//				panic("mock out the Update method")
//			},
//		}
//
//		// use mockedUseCase in code that requires UseCase
//		// and then make assertions.
//
//	}
type UseCaseMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, p *domain.Port) error

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, id string, upd Update) (*domain.Port, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P *domain.Port
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
			// Upd is the upd argument value.
			Upd Update
		}
	}
	lockCreate sync.RWMutex
	lockUpdate sync.RWMutex
}

// Create calls CreateFunc.
func (mock *UseCaseMock) Create(ctx context.Context, p *domain.Port) error {
	if mock.CreateFunc == nil {
		panic("UseCaseMock.CreateFunc: method is nil but UseCase.Create was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   *domain.Port
	}{
		Ctx: ctx,
		P:   p,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, p)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedUseCase.CreateCalls())
func (mock *UseCaseMock) CreateCalls() []struct {
	Ctx context.Context
	P   *domain.Port
} {
	var calls []struct {
		Ctx context.Context
		P   *domain.Port
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *UseCaseMock) Update(ctx context.Context, id string, upd Update) (*domain.Port, error) {
	if mock.UpdateFunc == nil {
		panic("UseCaseMock.UpdateFunc: method is nil but UseCase.Update was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
		Upd Update
	}{
		Ctx: ctx,
		ID:  id,
		Upd: upd,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, id, upd)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//
//	len(mockedUseCase.UpdateCalls())
func (mock *UseCaseMock) UpdateCalls() []struct {
	Ctx context.Context
	ID  string
	Upd Update
} {
	var calls []struct {
		Ctx context.Context
		ID  string
		Upd Update
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
