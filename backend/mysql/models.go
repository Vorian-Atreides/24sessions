package mysql

import (
	"github.com/Vorian-Atreides/24sessions/backend"
)

// Geolocation modelise the Geolocation for the MySQL layer
type Geolocation struct {
	IP      string `db:"ip"`
	City    string `db:"city"`
	Country string `db:"country"`
}

// FromModel copy the data from the backend model
func (g *Geolocation) FromModel(geo *backend.Geolocation) {
	g.IP = geo.IP
	g.City = geo.City
	g.Country = geo.Country
}

// ToModel instantiate a new backend model from the information stored in the
// instance
func (g *Geolocation) ToModel() *backend.Geolocation {
	return &backend.Geolocation{
		IP:      g.IP,
		City:    g.City,
		Country: g.Country,
	}
}
