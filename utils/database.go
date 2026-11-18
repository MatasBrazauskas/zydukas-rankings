package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DbAddress   string
	MaxOpenCons int
	MaxIdleCons int
	MaxIdleTime int
}

func getDatabaseConfig() (*DatabaseConfig, error) {

	dbAddress, dbAErr := EnvString("DB_ADDRESS")
	if dbAErr != nil {
		return nil, dbAErr
	}

	maxOpenCons, maxErr := EnvInt("DB_MAX_OPEN_CONS")
	if maxErr != nil {
		return nil, maxErr
	}

	maxIdleCons, iErr := EnvInt("DB_MAX_IDLE_CONS")
	if iErr != nil {
		return nil, iErr
	}

	maxIdleTime, tErr := EnvInt("DB_MAX_IDLE_TIME")
	if tErr != nil {
		return nil, tErr
	}

	return &DatabaseConfig{
		DbAddress:   dbAddress,
		MaxOpenCons: maxOpenCons,
		MaxIdleCons: maxIdleCons,
		MaxIdleTime: maxIdleTime,
	}, nil
}

func Dsn() (string, error) {
	dsnString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"

	host, hErr := EnvString("DB_HOST")
	if hErr != nil {
		return "", hErr
	}

	portNum, poErr := EnvInt("DB_PORT")
	if poErr != nil {
		return "", poErr
	}
	port := strconv.Itoa(portNum)

	user, uErr := EnvString("DB_USER")
	if uErr != nil {
		return "", uErr
	}

	password, pErr := EnvString("DB_PASSWORD")
	if pErr != nil {
		return "", pErr
	}

	dbName, nErr := EnvString("DB_NAME")
	if nErr != nil {
		return "", nErr
	}

	sslMode, sErr := EnvString("DB_SSL_MODE")
	if sErr != nil {
		return "", sErr
	}

	timezone, tErr := EnvString("DB_TIMEZONE")
	if tErr != nil {
		return "", tErr
	}

	dsn := fmt.Sprintf(dsnString, host, user, password, dbName, port, sslMode, timezone)

	return dsn, nil
}

func InitDatabase() (*gorm.DB, error) {
	dbConfig, dbErr := getDatabaseConfig()
	if dbErr != nil {
		return nil, dbErr
	}

	gormConfig := gorm.Config{
		TranslateError: true,
	}

	dsn, dsnErr := Dsn()
	if dsnErr != nil {
		return nil, dsnErr
	}

	db, err := gorm.Open(postgres.Open(dsn), &gormConfig)

	if err != nil {
		return nil, err
	}

	sqlDb, sErr := db.DB()
	if sErr != nil {
		return nil, sErr
	}

	sqlDb.SetMaxOpenConns(dbConfig.MaxOpenCons)
	sqlDb.SetMaxIdleConns(dbConfig.MaxIdleCons)
	sqlDb.SetConnMaxIdleTime(time.Duration(dbConfig.MaxIdleTime) * time.Second)

	return db, nil
}

func CloseDatabase(db *gorm.DB) {
	sqlDb, sqlErr := db.DB()
	if sqlErr != nil {
		log.Println(sqlErr)
		return
	}

	sqlDb.Close()
}
