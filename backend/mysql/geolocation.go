package mysql

import (
	"github.com/Vorian-Atreides/24sessions/backend"
)

const (
	CreateGeolocation Query = 100 + iota
	GetGeolocationByIP
)

var geolocationQueries = map[Query]string{
	CreateGeolocation: `
		INSERT INTO test.geolocations(ip, city, country)
		VALUES(:ip, :city, :country)
	`,
	GetGeolocationByIP: `
		SELECT * FROM test.geolocations
		WHERE ip = :ip
	`,
}

func (m *MySQL) CreateGeolocation(location *backend.Geolocation) error {
	query := &Geolocation{}
	query.FromModel(location)
	_, err := m.Stmts[CreateGeolocation].Exec(query)
	return err
}

func (m *MySQL) GetGeolocationByIP(ip string) (*backend.Geolocation, error) {
	location := &Geolocation{}
	query := &Geolocation{IP: ip}
	err := m.Stmts[GetGeolocationByIP].Get(location, query)
	return location.ToModel(), err
}
