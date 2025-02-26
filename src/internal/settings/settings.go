package settings

import (
	"github.com/caarlos0/env/v6"
)

var settingsSingleton *Settings

type Settings struct {
	MySQLHost            string `env:"MYSQL_HOST"`
	MySQLPort            int    `env:"MYSQL_PORT"`
	MySQLUser            string `env:"MYSQL_USER"`
	MySQLPassword        string `env:"MYSQL_PASSWORD"`
	MySQLDBName          string `env:"MYSQL_DBNAME"`
	MySQLSSLMode         string `env:"MYSQL_SSL_MODE"`
	MySQLTimeZone        string `env:"MYSQL_TIMEZONE"`
	MySQLMaxIdleConns    int    `env:"MYSQL_MAX_IDLE_CONNS"`
	MySQLMaxOpenConns    int    `env:"MYSQL_MAX_OPEN_CONNS"`
	MySQLConnMaxLifetime int    `env:"MYSQL_CONN_MAX_LIFETIME"`
}

func GetSettings() *Settings {
	if settingsSingleton != nil {
		return settingsSingleton
	}

	cfg := Settings{}
	env.Parse(&cfg)
	settingsSingleton = &cfg
	return settingsSingleton
}
