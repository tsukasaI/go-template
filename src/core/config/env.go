package config

import (
	env "github.com/caarlos0/env/v11"
)

type config struct {
	MySQLHost     string `env:"MYSQL_HOST"`
	MySQLUser     string `env:"MYSQL_USER"`
	MySQLPassword string `env:"MYSQL_PASSWORD"`
	MySQLPort     int    `env:"MYSQL_PORT" envDefault:"3306"`
	MySQLDatabase string `env:"MYSQL_DATABASE"`
}

var cfg config

func init() {
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
}

func MySQLHost() string {
	return cfg.MySQLHost
}

func MySQLUser() string {
	return cfg.MySQLUser
}

func MySQLPassword() string {
	return cfg.MySQLPassword
}

func MySQLPort() int {
	return cfg.MySQLPort
}

func MySQLDatabase() string {
	return cfg.MySQLDatabase
}
