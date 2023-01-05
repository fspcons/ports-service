package port_test

import (
	"context"
	"errors"
	"github.com/fspcons/ports-service/src/domain"
	"github.com/fspcons/ports-service/src/gateway/file"
	"github.com/fspcons/ports-service/src/gateway/ports"
	"github.com/fspcons/ports-service/src/usecases/port"
	"go.uber.org/zap"
	"testing"
)

// TODO If I had more time I'd add:
//  tests for the Update  method for both success and error cases mocking the DAO gateway

func TestCreate(t *testing.T) {
	logger, ctx := zap.NewNop(), context.TODO()
	fileError, portsError := errors.New("some file gateway error"), errors.New("some ports gateway error")

	scenarios := map[string]struct {
		portInput      *domain.Port
		expectedOutput error
		fileGateway    file.Gateway
		portsGateway   ports.Gateway
	}{
		"should be successful when no gateways return error": {
			portInput: &domain.Port{ID: "someID"}, expectedOutput: nil,
			fileGateway: &file.GatewayMock{
				CheckOnFileFunc: func(ctx context.Context, port *domain.Port) error {
					return nil
				},
			},
			portsGateway: &ports.GatewayMock{
				InsertFunc: func(ctx context.Context, rec *ports.Record) error {
					if rec.CreatedAt.IsZero() || rec.UpdatedAt.IsZero() {
						t.Error("metadata fields were not set on UseCase.Create")
					}

					return nil
				},
			},
		},
		"should fail when file gateway returns error": {
			portInput: &domain.Port{ID: "someID"}, expectedOutput: fileError,
			fileGateway: &file.GatewayMock{
				CheckOnFileFunc: func(ctx context.Context, port *domain.Port) error {
					return fileError
				},
			},
			portsGateway: &ports.GatewayMock{
				InsertFunc: func(ctx context.Context, rec *ports.Record) error {
					return nil
				},
			},
		},
		"should fail when ports gateway returns error": {
			portInput: &domain.Port{ID: "someID"}, expectedOutput: portsError,
			fileGateway: &file.GatewayMock{
				CheckOnFileFunc: func(ctx context.Context, port *domain.Port) error {
					return nil
				},
			},
			portsGateway: &ports.GatewayMock{
				InsertFunc: func(ctx context.Context, rec *ports.Record) error {
					return portsError
				},
			},
		},
	}

	t.Log("Given a list of scenarios for UseCase.Create")
	for scenario, v := range scenarios {
		t.Run(scenario, func(t *testing.T) {
			uc := port.NewUseCase(v.portsGateway, v.fileGateway, logger)

			result := uc.Create(ctx, v.portInput)
			if !errors.Is(result, v.expectedOutput) {
				t.Errorf("unexpected UseCase.Create result: expected %+v, got %+v", result, v.expectedOutput)
			}
		})
	}
}
