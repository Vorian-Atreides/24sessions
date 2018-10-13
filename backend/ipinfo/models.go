package ipinfo

import (
	"errors"

	"github.com/Vorian-Atreides/24sessions/backend"
)

type Geo struct {
	IP      string `json:"ip"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func (g *Geo) ToModel() *backend.Geolocation {
	return &backend.Geolocation{
		IP:      g.IP,
		City:    g.City,
		Country: g.Country,
	}
}

func (g *Geo) Validate() error {
	if len(g.City) > 0 && len(g.Country) > 0 && len(g.IP) > 0 {
		return nil
	}
	return errors.New("invalid geolocation")
}
