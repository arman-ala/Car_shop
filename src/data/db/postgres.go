package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/arman-ala/Car_shop/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDB(cfg config.Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=Asia/Tehran sslmode=%s", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBName, cfg.Postgres.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	dbClient = db
	err, sqlDB := PingDB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Postgres.MaxLifetime)

	log.Printf("Database connection was established successfully")
	return nil
}

func PingDB() (error, *sql.DB) {
	sqlDB, err := dbClient.DB()
	if err != nil {
		return err, nil
	}
	err = sqlDB.Ping()
	if err != nil {
		return err, nil
	}
	return nil, sqlDB
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB(db *sql.DB) {
	con, _ := dbClient.DB()
	con.Close()
}
