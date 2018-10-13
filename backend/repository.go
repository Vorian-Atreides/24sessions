package backend

type Geolocations interface {
	CreateGeolocation(location *Geolocation) error
	GetGeolocationByIP(ip string) (*Geolocation, error)
}

type Repository interface {
	Geolocations
}
