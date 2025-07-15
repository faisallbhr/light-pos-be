package database

import (
	"context"
	"log"
	"time"

	"github.com/faisallbhr/light-pos-be/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

var database *DB

func InitDB(cfg *config.Config) {
	dsn := cfg.GetDatabaseDSN()

	var logLevel logger.LogLevel
	if cfg.App.Environment == "production" {
		logLevel = logger.Silent
	} else {
		logLevel = logger.Info
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get SQL DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 30)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}

	database = &DB{db}
	log.Println("connected to MySQL with GORM:", cfg.Database.Name)
}

func GetDB() *DB {
	return database
}

func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// WithContext returns a new DB instance with context
func (db *DB) WithContext(ctx context.Context) *DB {
	return &DB{db.DB.WithContext(ctx)}
}

// Transaction wrapper with context
func (db *DB) WithTransaction(ctx context.Context, fn func(*gorm.DB) error) error {
	return db.WithContext(ctx).Transaction(fn)
}
