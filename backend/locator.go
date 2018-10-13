package backend

type Geolocator interface {
	FromIP(ip string) (*Geolocation, error)
}
