package database

import (
	"fmt"
	"log"
	"template-go-api/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbRead  *gorm.DB
	dbWrite *gorm.DB
	err     error
)

const MAX_IDLE_CONNECTIONS = 10
const MAX_OPEN_CONNECTIONS = 10
const MAX_LIFETIMES = time.Hour

func Init(cfg config.Config) {
	// Initialize read database connection
	if cfg.DBHostRead != "" {
		dsnRead := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
			cfg.DBHostRead, cfg.DBUser, cfg.DBPasswd, cfg.DBName, cfg.DBPort)

		dbRead, err = gorm.Open(postgres.Open(dsnRead), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err != nil {
			log.Fatalf("Failed to connect to read-only database: %v", err)
		}
	}

	// Initialize write database connection
	if cfg.DBHostWrite != "" {
		dsnWrite := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
			cfg.DBHostWrite, cfg.DBUser, cfg.DBPasswd, cfg.DBName, cfg.DBPort)

		dbWrite, err = gorm.Open(postgres.Open(dsnWrite), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to write database: %v", err)
		}
	}

	// If only one database is configured, use it for both read and write
	if dbRead == nil && dbWrite != nil {
		dbRead = dbWrite
	} else if dbWrite == nil && dbRead != nil {
		dbWrite = dbRead
	}

	// Setup connection pool
	if dbRead != nil || dbWrite != nil {
		dbPool()
		migrate()
	}
}

func dbPool() {
	if dbRead != nil {
		sqlDBRead, err := dbRead.DB()
		if err != nil {
			log.Fatalf("Failed to get database instance: %v", err)
		}
		sqlDBRead.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
		sqlDBRead.SetMaxOpenConns(MAX_OPEN_CONNECTIONS)
		sqlDBRead.SetConnMaxLifetime(MAX_LIFETIMES)
	}

	if dbWrite != nil {
		sqlDBWrite, err := dbWrite.DB()
		if err != nil {
			log.Fatalf("Failed to get database instance: %v", err)
		}
		sqlDBWrite.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
		sqlDBWrite.SetMaxOpenConns(MAX_OPEN_CONNECTIONS)
		sqlDBWrite.SetConnMaxLifetime(MAX_LIFETIMES)
	}
}

func migrate() {
	if dbWrite == nil {
		log.Println("No write database configured, skipping migration")
		return
	}

	fmt.Println("Starting database migration...")

	// Add your domain models here for auto migration
	// err = dbWrite.AutoMigrate(
	// 	&domains.User{},
	// 	&domains.Role{},
	// 	// Add other models here
	// )

	// if err != nil {
	// 	log.Fatalf("Failed to migrate: %v", err)
	// }

	fmt.Println("Database migration completed")
}

func GetReadDB() *gorm.DB {
	return dbRead
}

func GetWriteDB() *gorm.DB {
	return dbWrite
}
