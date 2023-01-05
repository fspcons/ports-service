package ports

import (
	"context"
	"go.uber.org/zap"
	"sync"
)

// Gateway allows fetch, creation and modification of ports records.
//
//go:generate moq -out gateway_mock.go . Gateway:GatewayMock
type Gateway interface {
	// Insert persists a record in the database.
	Insert(ctx context.Context, rec *Record) error
	// FindOneByID retrieves a record with a given id from the database.
	FindOneByID(ctx context.Context, id string) (*Record, error)
	// Update modifies an existing record in the database.
	Update(ctx context.Context, rec *Record) error
}

// NewInMemoryGateway builds a gateway to an in memory database.
func NewInMemoryGateway(logger *zap.Logger) Gateway {
	return &inMemoryDB{db: make(map[string]*Record, 0), lock: sync.Mutex{}, logger: logger}
}
