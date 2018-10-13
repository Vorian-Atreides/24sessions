package backend_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Vorian-Atreides/24sessions/backend/ipinfo"
	"github.com/stretchr/testify/assert"

	"github.com/Vorian-Atreides/24sessions/backend"
)

var (
	errNotFound = errors.New("IP not cached")
)

type GeolocatorMock struct {
	geolocation *backend.Geolocation
	err         error
}

func (g *GeolocatorMock) FromIP(ip string) (*backend.Geolocation, error) {
	return g.geolocation, g.err
}

type GeolocationRepoMock struct {
	items map[string]*backend.Geolocation
	err   error
}

func (g *GeolocationRepoMock) CreateGeolocation(location *backend.Geolocation) error {
	if g.err != nil {
		return g.err
	}
	g.items[location.IP] = location
	return nil
}

func (g *GeolocationRepoMock) GetGeolocationByIP(ip string) (*backend.Geolocation, error) {
	if g.err != nil {
		return nil, g.err
	}
	location, ok := g.items[ip]
	if !ok {
		return nil, errNotFound
	}
	return location, nil
}

type GetByIPRequestTest struct {
	repo       backend.Repository
	geolocator backend.Geolocator

	request  *backend.GetByIPRequest
	expected *backend.Response
}

var GetByIPRequestTests = map[string]GetByIPRequestTest{
	"not cached and local IP": {
		repo:       &GeolocationRepoMock{items: map[string]*backend.Geolocation{}},
		geolocator: &GeolocatorMock{err: ipinfo.ErrInvalidGeolocation},

		request: &backend.GetByIPRequest{
			IP: "192.168.0.30",
		},
		expected: &backend.Response{
			ContentType: "application/json",
			StatusCode:  http.StatusBadGateway,
			ErrMessage:  "geolocator_error",
			Err:         ipinfo.ErrInvalidGeolocation,
		},
	},
	"not cached and IP found": {
		repo: &GeolocationRepoMock{items: map[string]*backend.Geolocation{}},
		geolocator: &GeolocatorMock{geolocation: &backend.Geolocation{
			IP:      "8.8.8.8",
			City:    "Mountain View",
			Country: "US",
		}},

		request: &backend.GetByIPRequest{
			IP: "8.8.8.8",
		},
		expected: &backend.Response{
			ContentType: "application/json",
			StatusCode:  http.StatusOK,
			Data: &backend.Geolocation{
				IP:      "8.8.8.8",
				City:    "Mountain View",
				Country: "US",
			},
		},
	},
	"cached IP": {
		repo: &GeolocationRepoMock{items: map[string]*backend.Geolocation{
			"8.8.8.8": &backend.Geolocation{
				IP:      "8.8.8.8",
				City:    "Mountain View",
				Country: "US",
			},
		}},
		// Ensure that this piece of code will be unused in this flow
		geolocator: &GeolocatorMock{err: errNotFound},

		request: &backend.GetByIPRequest{
			IP: "8.8.8.8",
		},
		expected: &backend.Response{
			ContentType: "application/json",
			StatusCode:  http.StatusOK,
			Data: &backend.Geolocation{
				IP:      "8.8.8.8",
				City:    "Mountain View",
				Country: "US",
			},
		},
	},
}

func TestGetByIP(t *testing.T) {
	for key, test := range GetByIPRequestTests {
		controller := backend.NewGeolocationController(test.repo, test.geolocator)
		t.Run(key, func(t *testing.T) {
			result := controller.GetByIP(test.request)
			assert.Equal(t, test.expected, result)

			// Ensure the persistent layer has been called
			if test.expected.Err == nil {
				geolocation, err := test.repo.GetGeolocationByIP(test.request.IP)
				assert.NoError(t, err)
				assert.NotNil(t, geolocation)
			}
		})
	}
}
