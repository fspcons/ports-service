package port

import (
	"context"
	"github.com/fspcons/ports-service/src/domain"
	"github.com/fspcons/ports-service/src/gateway/file"
	"github.com/fspcons/ports-service/src/gateway/ports"
	"go.uber.org/zap"
	"time"
)

type uc struct {
	ports  ports.Gateway
	files  file.Gateway
	logger *zap.Logger
}

func (ref *uc) handle(err error) error {
	if err != nil {
		ref.logger.Error(err.Error())
	}
	return err
}

// Create produces a new domain.Port.
func (ref *uc) Create(ctx context.Context, port *domain.Port) error {
	if port == nil || !port.IsValid() {
		return ref.handle(domain.ErrInvalidPort)
	}

	if err := ref.files.CheckOnFile(ctx, port); err != nil {
		return ref.handle(err)
	}

	now := time.Now().UTC()
	rec := toRecord(port)
	rec.CreatedAt = now
	rec.UpdatedAt = now

	return ref.handle(ref.ports.Insert(ctx, rec))
}

// Update modifies a given domain.Port.
func (ref *uc) Update(ctx context.Context, id string, upd Update) (*domain.Port, error) {
	if err := upd.Validate(); err != nil {
		return nil, ref.handle(err)
	}

	oldPort, err := ref.ports.FindOneByID(ctx, id)
	if err != nil {
		return nil, ref.handle(err)
	}

	if upd.Name != nil {
		oldPort.Name = *upd.Name
	}
	if upd.City != nil {
		oldPort.City = *upd.City
	}
	if upd.Country != nil {
		oldPort.Country = *upd.Country
	}
	if upd.Alias != nil {
		oldPort.Alias = upd.Alias
	}
	if upd.Regions != nil {
		oldPort.Regions = upd.Regions
	}
	if upd.Coordinates != nil {
		oldPort.Coordinates = upd.Coordinates
	}
	if upd.Province != nil {
		oldPort.Province = *upd.Province
	}
	if upd.Timezone != nil {
		oldPort.Timezone = *upd.Timezone
	}
	if upd.Unlocs != nil {
		oldPort.Unlocs = upd.Unlocs
	}
	if upd.Code != nil {
		oldPort.Code = *upd.Code
	}

	oldPort.UpdatedAt = time.Now().UTC()

	return fromRecord(oldPort), ref.handle(ref.ports.Update(ctx, oldPort))
}
