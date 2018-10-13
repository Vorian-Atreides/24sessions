package ipinfo

import "github.com/urfave/cli"

// Config describe the required arguments to instantiate an IPInfo geolocator.
// The structure is compliant with caarlos0/env
type Config struct {
	// APIToken to allow the authenticate the application with IpInfo
	APIToken string `json:"api_token" env:"API_TOKEN"`
}

// GetParams compliant method with urfave/cli's Flags
func (c *Config) GetParams() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "api_token",
			Usage:       "APIToken to allow the authenticate the application with IpInfo",
			EnvVar:      "API_TOKEN",
			Destination: &c.APIToken,
		},
	}
}
