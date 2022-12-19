package ports

import (
	"context"
	"github.com/fspcons/ports-service/src/domain"
	"go.uber.org/zap"
	"sync"
)

type inMemoryDB struct {
	db     map[string]*Record
	lock   sync.Mutex
	logger *zap.Logger
}

// Insert persists a record in the database.
func (ref *inMemoryDB) Insert(ctx context.Context, rec *Record) error {
	ref.lock.Lock()
	defer ref.lock.Unlock()

	if _, err := ref.FindOneByID(ctx, rec.ID); err == domain.ErrNoRecords {
		ref.db[rec.ID] = rec
		return nil
	}

	return domain.ErrRecordAlreadyExists //in case idempotency was necessary we could return nil here to respond success on the request
}

// FindOneByID retrieves a record with a given id from the database.
func (ref *inMemoryDB) FindOneByID(_ context.Context, id string) (*Record, error) {
	rec, exists := ref.db[id]
	if !exists {
		return nil, domain.ErrNoRecords
	}

	return rec, nil
}

// Update modifies an existing record in the database.
func (ref *inMemoryDB) Update(_ context.Context, rec *Record) error {
	ref.lock.Lock()
	defer ref.lock.Unlock()

	if _, exists := ref.db[rec.ID]; !exists {
		return domain.ErrNoRecords
	}

	ref.db[rec.ID] = rec

	return nil
}
