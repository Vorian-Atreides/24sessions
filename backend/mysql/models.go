package mysql

import (
	"github.com/Vorian-Atreides/24sessions/backend"
)

type Geolocation struct {
	IP      string `db:"ip"`
	City    string `db:"city"`
	Country string `db:"country"`
}

func (g *Geolocation) FromModel(geo *backend.Geolocation) {
	g.IP = geo.IP
	g.City = geo.City
	g.Country = geo.Country
}

func (g *Geolocation) ToModel() *backend.Geolocation {
	return &backend.Geolocation{
		IP:      g.IP,
		City:    g.City,
		Country: g.Country,
	}
}
