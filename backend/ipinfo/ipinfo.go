package ipinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Vorian-Atreides/24sessions/backend"
)

const (
	base = "https://ipinfo.io"
)

// IPInfo implement the Geolocator interface with ipinfo.io
type IPInfo struct {
	apiToken string
}

// New instantiate a new IPInfo
func New(cfg *Config) *IPInfo {
	return &IPInfo{
		apiToken: cfg.APIToken,
	}
}

// FromIP retrieve the geolocation from an IP
func (i *IPInfo) FromIP(ip string) (*backend.Geolocation, error) {
	// We use the query parameter method for the token by convenience,
	// we should replace it with the header if we implement other endpoints.
	url := fmt.Sprintf("%s/%s/geo?token=%s", base, ip, i.apiToken)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	location := &Geo{}
	if err := json.Unmarshal(bytes, location); err != nil {
		return nil, err
	}
	return location.ToModel(), location.Validate()
}
