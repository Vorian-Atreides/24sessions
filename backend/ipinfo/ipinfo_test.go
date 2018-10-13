package ipinfo_test

import (
	"testing"

	"github.com/Vorian-Atreides/24sessions/backend"
	"github.com/Vorian-Atreides/24sessions/backend/ipinfo"
	"github.com/stretchr/testify/assert"
)

var cfg = &ipinfo.Config{
	APIToken: "00dfa9010b2e69",
}

type FromIPTest struct {
	ip       string
	expected *backend.Geolocation
	err      error
}

var FromIPTests = map[string]FromIPTest{
	"Google DNS": {
		ip:       "8.8.8.8",
		expected: &backend.Geolocation{IP: "8.8.8.8", City: "Mountain View", Country: "US"},
	},
	"Romorantin": {
		ip:       "78.238.100.100",
		expected: &backend.Geolocation{IP: "78.238.100.100", City: "Romorantin", Country: "FR"},
	},
	"local network": {
		ip: "192.168.0.30",
		expected: &backend.Geolocation{
			IP: "192.168.0.30",
		},
		err: ipinfo.ErrInvalidGeolocation,
	},
}

func TestFromIP(t *testing.T) {
	geolocator := ipinfo.New(cfg)

	for key, test := range FromIPTests {
		t.Run(key, func(t *testing.T) {
			result, err := geolocator.FromIP(test.ip)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.err, err)
		})
	}
}
