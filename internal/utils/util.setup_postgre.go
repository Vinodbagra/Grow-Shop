package utils

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/snykk/grow-shop/internal/config"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/drivers"
)

func SetupPostgresConnection() (*sqlx.DB, error) {
	var dsn string
	switch config.AppConfig.Environment {
	case constants.EnvironmentDevelopment:
		dsn = config.AppConfig.DBPostgreDsn
	case constants.EnvironmentProduction:
		dsn = config.AppConfig.DBPostgreURL
	}

	// Setup sqlx config of postgreSQL
	config := drivers.SQLXConfig{
		DriverName:     config.AppConfig.DBPostgreDriver,
		DataSourceName: dsn,
		MaxOpenConns:   100,
		MaxIdleConns:   10,
		MaxLifetime:    15 * time.Minute,
	}

	// Initialize postgreSQL connection with sqlx
	conn, err := config.InitializeSQLXDatabase()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
