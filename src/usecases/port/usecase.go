package port

import (
	"context"
	"github.com/fspcons/ports-service/src/domain"
	"github.com/fspcons/ports-service/src/gateway/file"
	"github.com/fspcons/ports-service/src/gateway/ports"
	"go.uber.org/zap"
)

// UseCase allows fetch, creation and modification of domain.Port.
//
//go:generate moq -out usecase_mock.go . UseCase:UseCaseMock
type UseCase interface {
	// Create produces a new domain.Port.
	Create(ctx context.Context, p *domain.Port) error
	// Update modifies a given domain.Port.
	Update(ctx context.Context, id string, upd Update) (*domain.Port, error)
}

// NewUseCase returns a new Port useCase.
func NewUseCase(pg ports.Gateway, files file.Gateway, logger *zap.Logger) UseCase {
	return &uc{ports: pg, files: files, logger: logger}
}
