package backend

// Geolocations interface the persistent layer for the Geolocation
// resources
type Geolocations interface {
	CreateGeolocation(location *Geolocation) error
	GetGeolocationByIP(ip string) (*Geolocation, error)
}

// Repository interface the whole persistent layer
type Repository interface {
	Geolocations
}
