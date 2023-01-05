package domain

import "github.com/fspcons/ports-service/src/utils"

type Port struct {
	ID          string    `json:"id,omitempty" db:"id"`
	Name        string    `json:"name,omitempty" db:"name"`
	City        string    `json:"city,omitempty" db:"city"`
	Country     string    `json:"country,omitempty" db:"country"`
	Alias       []string  `json:"alias,omitempty" db:"alias"`
	Regions     []string  `json:"regions,omitempty" db:"regions"`
	Coordinates []float32 `json:"coordinates,omitempty" db:"coordinates"`
	Province    string    `json:"province,omitempty" db:"province"`
	Timezone    string    `json:"timezone,omitempty" db:"timezone"`
	Unlocs      []string  `json:"unlocs,omitempty" db:"unlocs"`
	Code        string    `json:"code,omitempty" db:"code"`
}

func (ref Port) IsValid() bool {
	//Port validation logic here...
	return !utils.IsEmpty(ref.ID)
}
