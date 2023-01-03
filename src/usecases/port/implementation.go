package port

import (
	"bufio"
	"context"
	"github.com/fspcons/ports-service/src/domain"
	"github.com/fspcons/ports-service/src/gateway/ports"
	"go.uber.org/zap"
	"io"
	"os"
	"sync"
	"time"
)

type uc struct {
	ports        ports.Gateway
	logger       *zap.Logger
	portFilePath string
}

func (ref *uc) handle(err error) error {
	if err != nil {
		ref.logger.Error(err.Error())
	}
	return err
}

// CheckOnFile if the port record exists on the json file
func (ref *uc) CheckOnFile(_ context.Context, _ *domain.Port) error {
	f, err := os.Open(ref.portFilePath)
	if err != nil {
		ref.logger.Error("failed to open ports file", zap.Error(err))
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			ref.logger.Warn("failed to close ports file", zap.Error(err))
		}
	}()

	//sync pools to reuse the memory and decrease the pressure on the GC
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 100*1024) //chunk size
		return lines
	}}

	r := bufio.NewReader(f)
	for {
		buf := linesPool.Get().([]byte) //reading the file by chunks
		n, err := r.Read(buf)
		buf = buf[:n]
		_ = buf //avoid lint errors
		if n == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				ref.logger.Error("failed to read a file chunk into the buffer", zap.Error(err))
				break
			}
			return err
		}
	}

	//TODO From here forward I'd most likely run a Regex check or a similar method, trying to find the port ID as the KEY
	// among the json file records. I believe this would perform better than trying to unmarshall the json records from the file.
	// In case I find it I'd check to see if the user provided data matches and would return a NIL error.
	// Otherwise I'd return some error saying that the Port record provided is invalid or something of the sorts.
	// Additionally if I could assume the keys on the file are alphabetical I could use a binary search method instead of a linear one,
	// and I could also suggest a file partitioning system based on the keys so I could have smaller files and a smarter search mechanic in that case.

	return nil
}

// Create produces a new domain.Port.
func (ref *uc) Create(ctx context.Context, port *domain.Port) error {
	if port == nil || !port.IsValid() {
		return ref.handle(domain.ErrInvalidPort)
	}

	if err := ref.CheckOnFile(ctx, port); err != nil {
		return ref.handle(domain.ErrInvalidPort)
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
		return nil, err
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
