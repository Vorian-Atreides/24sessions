package backend

// Geolocation couple an IP address and its geolocation
type Geolocation struct {
	IP      string `json:"ip"`
	City    string `json:"city"`
	Country string `json:"country"`
}
