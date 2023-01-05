package ports_test

import (
	"context"
	"errors"
	"github.com/fspcons/ports-service/src/domain"
	"github.com/fspcons/ports-service/src/gateway/ports"
	"go.uber.org/zap"
	"testing"
)

func TestInsert(t *testing.T) {
	gateway, ctx := ports.NewInMemoryGateway(zap.NewNop()), context.TODO()

	scenarios := map[string]struct {
		recordInput    *ports.Record
		expectedOutput error
	}{
		"should be successful for a non existing ID": {recordInput: &ports.Record{
			Port: domain.Port{ID: "someID"},
		}, expectedOutput: nil},
		"should fail for an existing ID": {recordInput: &ports.Record{
			Port: domain.Port{ID: "someID"},
		}, expectedOutput: domain.ErrRecordAlreadyExists},
	}

	t.Log("Given a list of scenarios for Gateway.Insert")
	for scenario, v := range scenarios {
		t.Run(scenario, func(t *testing.T) {
			result := gateway.Insert(ctx, v.recordInput)
			if !errors.Is(result, v.expectedOutput) {
				t.Errorf("unexpected Gateway.Insert result: expected %+v, got %+v", result, v.expectedOutput)
			}
		})
	}
}

func TestFindOneByID(t *testing.T) {
	gateway, ctx := ports.NewInMemoryGateway(zap.NewNop()), context.TODO()

	existingRecord := &ports.Record{Port: domain.Port{ID: "someID"}}
	if err := gateway.Insert(ctx, existingRecord); err != nil {
		t.Fatal(err.Error())
	}

	scenarios := map[string]struct {
		idInput        string
		expectedOutput *ports.Record
		expectedErr    error
	}{
		"should be successful for an existing ID": {idInput: existingRecord.ID, expectedOutput: existingRecord, expectedErr: nil},
		"should fail for a non existing ID":       {idInput: "someOtherID", expectedOutput: nil, expectedErr: domain.ErrNoRecords},
	}

	t.Log("Given a list of scenarios for Gateway.FindOneByID")
	for scenario, v := range scenarios {
		t.Run(scenario, func(t *testing.T) {
			result, err := gateway.FindOneByID(ctx, v.idInput)
			if !errors.Is(err, v.expectedErr) {
				t.Errorf("unexpected Gateway.FindOneByID error output: expected %+v, got %+v", err, v.expectedErr)
			}
			if result != nil && v.expectedOutput != nil {
				if result.ID != v.expectedOutput.ID {
					t.Errorf("unexpected Gateway.FindOneByID result: expected %s, got %s", result.ID, v.expectedOutput.ID)
				}
			} else if result != v.expectedOutput {
				t.Errorf("mismatching Gateway.FindOneByID result: expected %+v, got %+v", result, v.expectedOutput)
			}
		})
	}
}

// TODO If I had more time I'd add:
//  tests for Update methods
//  for both success and error cases.
