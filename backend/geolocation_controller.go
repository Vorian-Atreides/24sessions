package backend

import (
	"net/http"

	"github.com/sirupsen/logrus"

	validator "gopkg.in/go-playground/validator.v9"
)

// GeolocationController implements the endpoints related to the geolocation
type GeolocationController struct {
	validator *validator.Validate

	repo       Geolocations
	geolocator Geolocator
}

// NewGeolocationController instantiate a valid GeolocationController
func NewGeolocationController(repo Repository, geolocator Geolocator) *GeolocationController {
	return &GeolocationController{
		validator:  validator.New(),
		repo:       repo,
		geolocator: geolocator,
	}
}

// GetByIPRequest is the input for the GetByIP endpoint
type GetByIPRequest struct {
	IP string `json:"-" validate:"required,ip"`
}

// GetByIP will lookup in the persistent layer if the IP has already been
// searched, otherwise it will ask the Geolocator to retrieve the geolocation
// and save it for the next usage.
func (g *GeolocationController) GetByIP(req *GetByIPRequest) *Response {
	if err := g.validator.Struct(req); err != nil {
		return NewResponse().WithErr(err, "invalid_parameters").
			WithStatusCode(http.StatusBadRequest)
	}

	// First check if the IP is stored in the persistent layer
	location, err := g.repo.GetGeolocationByIP(req.IP)
	if err != nil {
		logrus.WithError(err).Warn("location_not_found")
		// If it isn't, ask the locator
		location, err = g.geolocator.FromIP(req.IP)
		if err != nil {
			return NewResponse().WithErr(err, "geolocator_error").
				WithStatusCode(http.StatusBadGateway)
		}
		// Store the freshly queried location in the persistent layer
		if err := g.repo.CreateGeolocation(location); err != nil {
			logrus.WithError(err).Warn("repo_create_geolocation")
		}
	}
	return NewResponse().WithStatusCode(http.StatusOK).WithData(location)
}
