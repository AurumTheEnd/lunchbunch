package database

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/joho/godotenv"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
)

func CreateDbIfNotExists() (err error) {
	var myEnv map[string]string
	myEnv, err = godotenv.Read()

	var tempDsn = url.URL{
		User:   url.UserPassword(myEnv["POSTGRES_USER"], myEnv["POSTGRES_PASSWORD"]),
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%s", myEnv["POSTGRES_HOST"], myEnv["POSTGRES_PORT"]),
	}
	tempDb, err := gorm.Open(postgres.Open(tempDsn.String()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	//
	//_ = tempDb.Exec("DROP DATABASE " + myEnv["POSTGRES_DATABASE"] + ";").Commit()
	tempDb = tempDb.Exec("CREATE DATABASE " + myEnv["POSTGRES_DATABASE"] + ";").Commit()

	// ignore everything apart from duplicate_database error (https://www.postgresql.org/docs/current/errcodes-appendix.html)
	var pgErr *pgconn.PgError
	if errors.As(tempDb.Error, &pgErr) {
		if pgErr.Code != DatabaseAlreadyExistsErrorCode {
			err = tempDb.Error
		}
	}

	return
}

func Setup() (db *gorm.DB, err error) {
	var myEnv map[string]string
	myEnv, err = godotenv.Read()

	var dsn = url.URL{
		User:     url.UserPassword(myEnv["POSTGRES_USER"], myEnv["POSTGRES_PASSWORD"]),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", myEnv["POSTGRES_HOST"], myEnv["POSTGRES_PORT"]),
		Path:     myEnv["POSTGRES_DATABASE"],
		RawQuery: (&url.Values{"sslmode": []string{"disable"}, "timezone": []string{myEnv["POSTGRES_TIMEZONE"]}}).Encode(),
	}
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn.String(),
		PreferSimpleProtocol: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	err = db.AutoMigrate(&models.RestaurantSnapshot{}, &models.Restaurant{}, &models.MenuItem{}, &models.User{})

	return
}
