package utils

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DatabaseConfig struct {
	dbAddress   string
	maxOpenCons int
	maxIdleCons int
	maxIdleTime string
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		dbAddress:   EnvString("DB_ADDRESS", DbAddress),
		maxOpenCons: EnvInt("DB_MAX_OPEN_CONS", MaxOpenCons),
		maxIdleCons: EnvInt("DB_MAX_IDLE_CONS", MaxIdleCons),
		maxIdleTime: EnvString("DB_MAX_IDLE_TIME", MaxIdleTime),
	}
}

func InitDatabase(dbConfig *DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConfig.dbAddress)
	if err != nil {
		log.Println("Error connecting to database")
		return nil, err
	}

	db.SetMaxOpenConns(dbConfig.maxOpenCons)
	db.SetMaxIdleConns(dbConfig.maxIdleCons)

	duration, e := time.ParseDuration(dbConfig.maxIdleTime)
	if e != nil {
		return nil, e
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDatabase(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}
