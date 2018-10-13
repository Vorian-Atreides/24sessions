package main

import (
	"github.com/Vorian-Atreides/24sessions/backend/ipinfo"
	"github.com/Vorian-Atreides/24sessions/backend/mysql"
	"github.com/urfave/cli"
)

// Config define the configuration to run the application
type Config struct {
	IPInfoConfig ipinfo.Config `json:"ip_info"`
	MySQLConfig  mysql.Config  `json:"mysql"`
}

// GetParams integrate AppConfig with cli library
func (c *Config) GetParams() []cli.Flag {
	params := []cli.Flag{}
	params = append(params, c.IPInfoConfig.GetParams()...)
	params = append(params, c.MySQLConfig.GetParams()...)
	return params
}
