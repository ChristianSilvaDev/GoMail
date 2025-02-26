package container

import (
	"time"
	"github.com/ChristianSilvaDev/GoMail/src/internal/infra"
	"github.com/ChristianSilvaDev/GoMail/src/internal/settings"
)

func ProvideMySQLDatabase() *infra.MySQLDatabase {
	settings := settings.GetSettings()
	return &infra.MySQLDatabase{
		Host: settings.MySQLHost,
		Port: settings.MySQLPort,
		User: settings.MySQLUser,
		Password: settings.MySQLPassword,
		DBName: settings.MySQLDBName,
		SSLMode: settings.MySQLSSLMode,
		TimeZone: settings.MySQLTimeZone,
		MaxIdleConns: settings.MySQLMaxIdleConns,
		MaxOpenConns: settings.MySQLMaxIdleConns,
		ConnMaxLifetime: time.Duration(settings.MySQLConnMaxLifetime) * time.Hour,
	}
}
