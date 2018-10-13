package ipinfo

import (
	"errors"

	"github.com/Vorian-Atreides/24sessions/backend"
)

var (
	// ErrInvalidGeolocation raised when ipinfo doesn't successfully retrieve
	// the geolocation
	ErrInvalidGeolocation = errors.New("invalid geolocation")
)

// Geo response from ipinfo.io /geo
type Geo struct {
	IP      string `json:"ip"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// ToModel copy the instance's data to a new backend model
func (g *Geo) ToModel() *backend.Geolocation {
	return &backend.Geolocation{
		IP:      g.IP,
		City:    g.City,
		Country: g.Country,
	}
}

// Validate if the retrieve information is valid
func (g *Geo) Validate() error {
	if len(g.City) > 0 && len(g.Country) > 0 && len(g.IP) > 0 {
		return nil
	}
	return ErrInvalidGeolocation
}
