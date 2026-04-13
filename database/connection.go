package database

import (
	"fmt"
	"log"
	"golfscoreid-jng/config"
	"golfscoreid-jng/domains"
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

	err = dbWrite.AutoMigrate(
		&domains.User{},
		&domains.Player{},
		&domains.Hole{},
		&domains.Score{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	fmt.Println("Database migration completed")

	// Seed Jatinangor Golf Course hole data if not already seeded
	seedHoles()
}

func seedHoles() {
	var count int64
	dbWrite.Model(&domains.Hole{}).Count(&count)
	if count > 0 {
		return // Already seeded
	}

	fmt.Println("Seeding Jatinangor Golf Course holes...")

	holes := []domains.Hole{
		{HoleNumber: 1, Par: 4, HCP: 15, Black: 430, Blue: 398, White: 376, Red: 352},
		{HoleNumber: 2, Par: 5, HCP: 18, Black: 554, Blue: 522, White: 509, Red: 492},
		{HoleNumber: 3, Par: 4, HCP: 2, Black: 402, Blue: 381, White: 360, Red: 335},
		{HoleNumber: 4, Par: 3, HCP: 12, Black: 176, Blue: 154, White: 136, Red: 121},
		{HoleNumber: 5, Par: 5, HCP: 14, Black: 516, Blue: 505, White: 478, Red: 455},
		{HoleNumber: 6, Par: 3, HCP: 10, Black: 190, Blue: 175, White: 154, Red: 133},
		{HoleNumber: 7, Par: 4, HCP: 6, Black: 418, Blue: 392, White: 365, Red: 332},
		{HoleNumber: 8, Par: 3, HCP: 8, Black: 212, Blue: 186, White: 140, Red: 113},
		{HoleNumber: 9, Par: 5, HCP: 16, Black: 533, Blue: 502, White: 479, Red: 451},
		{HoleNumber: 10, Par: 4, HCP: 13, Black: 407, Blue: 381, White: 360, Red: 338},
		{HoleNumber: 11, Par: 5, HCP: 17, Black: 573, Blue: 541, White: 516, Red: 494},
		{HoleNumber: 12, Par: 3, HCP: 15, Black: 210, Blue: 190, White: 164, Red: 140},
		{HoleNumber: 13, Par: 4, HCP: 3, Black: 441, Blue: 410, White: 389, Red: 356},
		{HoleNumber: 14, Par: 4, HCP: 1, Black: 463, Blue: 432, White: 410, Red: 380},
		{HoleNumber: 15, Par: 3, HCP: 11, Black: 208, Blue: 189, White: 164, Red: 125},
		{HoleNumber: 16, Par: 4, HCP: 7, Black: 421, Blue: 398, White: 372, Red: 341},
		{HoleNumber: 17, Par: 3, HCP: 4, Black: 185, Blue: 168, White: 150, Red: 132},
		{HoleNumber: 18, Par: 5, HCP: 9, Black: 548, Blue: 521, White: 498, Red: 470},
	}

	if err := dbWrite.Create(&holes).Error; err != nil {
		log.Printf("Failed to seed holes: %v", err)
		return
	}

	fmt.Println("Jatinangor Golf Course holes seeded successfully")
}

func GetReadDB() *gorm.DB {
	return dbRead
}

func GetWriteDB() *gorm.DB {
	return dbWrite
}
