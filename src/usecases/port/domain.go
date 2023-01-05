package port

import (
	"github.com/fspcons/ports-service/src/domain"
	"github.com/fspcons/ports-service/src/gateway/ports"
)

// Update domain.Port update payload.
type Update struct {
	Name        *string   `json:"name,omitempty"`
	City        *string   `json:"city,omitempty"`
	Country     *string   `json:"country,omitempty"`
	Alias       []string  `json:"alias,omitempty"`
	Regions     []string  `json:"regions,omitempty"`
	Coordinates []float32 `json:"coordinates,omitempty"`
	Province    *string   `json:"province,omitempty"`
	Timezone    *string   `json:"timezone,omitempty"`
	Unlocs      []string  `json:"unlocs,omitempty"`
	Code        *string   `json:"code,omitempty"`
}

func (ref Update) Validate() error {
	//Here would go update payload validation code ...
	return nil
}

// toRecord parses the Port to a ports.Record.
func toRecord(p *domain.Port) *ports.Record {
	return &ports.Record{
		Port: domain.Port{
			ID:          p.ID,
			Name:        p.Name,
			City:        p.City,
			Country:     p.Country,
			Alias:       p.Alias,
			Regions:     p.Regions,
			Coordinates: p.Coordinates,
			Province:    p.Province,
			Timezone:    p.Timezone,
		},
	}
}

// fromRecord parses a record to a domain.Port.
func fromRecord(p *ports.Record) *domain.Port {
	return &p.Port
}
