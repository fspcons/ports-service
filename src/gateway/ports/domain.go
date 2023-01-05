package ports

import (
	"github.com/fspcons/ports-service/src/domain"
	"time"
)

// Record represents a port record.
type Record struct {
	domain.Port
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}
