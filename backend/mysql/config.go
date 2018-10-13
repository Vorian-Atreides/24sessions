package mysql

import (
	"fmt"

	"github.com/urfave/cli"
)

// Config describe the required arguments used to open an SQL connection
// The structure is compliant with caarlos0/env
type Config struct {
	// Host of the instance
	Host string `json:"host" env:"DB_HOST" envDefault:"localhost"`
	// UserName of the DB's user
	UserName string `json:"user_name" env:"DB_USERNAME" envDefault:"dev"`
	// Password of the DB's user
	Password string `json:"password" env:"DB_PASSWORD" envDefault:"password"`
	// Database to access
	Database string `json:"database" env:"DB_NAME" envDefault:"rooster"`
	// Certificate used to encrypt the connection
	Certificate string `json:"certificate" env:"DB_CERTIFICATE"`
}

// GetParams compliant method with urfave/cli's Flags
func (c *Config) GetParams() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "db_username",
			Usage:       "Username for the dabatase connection",
			EnvVar:      "DB_USERNAME",
			Value:       "test_user",
			Destination: &c.UserName,
		},
		cli.StringFlag{
			Name:        "db_password",
			Usage:       "Password for the database connection",
			EnvVar:      "DB_PASSWORD",
			Value:       "secret",
			Destination: &c.Password,
		},
		cli.StringFlag{
			Name:        "db_host",
			Usage:       "Host where is located the database",
			EnvVar:      "DB_HOST",
			Value:       "localhost",
			Destination: &c.Host,
		},
		cli.StringFlag{
			Name:        "db_name",
			Usage:       "Database name used to store the data",
			EnvVar:      "DB_NAME",
			Value:       "test",
			Destination: &c.Database,
		},
	}
}

// MySQLConnectionString produce a connection string for MySQL
func (c *Config) MySQLConnectionString() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?collation=utf8_general_ci&parseTime=true",
		c.UserName, c.Password, c.Host, c.Database)
	return dsn
}
