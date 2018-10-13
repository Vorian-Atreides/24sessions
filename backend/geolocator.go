package backend

// Geolocator interface to retrieve geolocation
type Geolocator interface {
	FromIP(ip string) (*Geolocation, error)
}
