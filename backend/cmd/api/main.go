package main

import (
	"net/http"
	"os"

	"github.com/Vorian-Atreides/24sessions/backend"
	"github.com/Vorian-Atreides/24sessions/backend/ipinfo"
	"github.com/Vorian-Atreides/24sessions/backend/mysql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	version = "0.1.0"
	name    = "api"
	usage   = "24sessions assessment"
)

func run(config *Config) cli.ActionFunc {
	return func(c *cli.Context) error {
		geolocator := ipinfo.New(&config.IPInfoConfig)
		repo, err := mysql.New(&config.MySQLConfig)
		if err != nil {
			return err
		}
		defer repo.Close()

		geoController := geoController{
			controller: backend.NewGeolocationController(repo, geolocator),
		}
		// Route the API, to externalise if the API grow
		engine := gin.Default()
		engine.Use(cors.Default())
		engine.Handle(http.MethodGet, "/geolocation/:ip", geoController.getByIP)
		return engine.Run()
	}
}

func main() {
	config := &Config{}
	app := cli.NewApp()
	app.Version = version
	app.Name = name
	app.Usage = usage
	app.Flags = config.GetParams()
	app.Action = run(config)
	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Fatal("application_main")
	}
}
