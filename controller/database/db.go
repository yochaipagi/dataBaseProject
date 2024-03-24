package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dsnKey = "DSN"
)

var DB *gorm.DB

func SetupDB() error {
	dsn := os.Getenv(dsnKey)
	if dsn == "" {
		return fmt.Errorf("environment variable %s is not set", dsnKey)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	DB = db

	err = migrateDB()
	if err != nil {
		return fmt.Errorf("failed to migrate the database: %v", err)
	}

	if isDatabasePopulated() {
		return nil
	}

	err = populateDB()
	if err != nil {
		return fmt.Errorf("failed to populate the database: %v", err)
	}

	return nil
}

func migrateDB() error {
	err := DB.AutoMigrate(
		&Article{},
		&ArticleLine{},
		&ArticleWord{},
		&WordGroup{},
		&Word{},
		&LinguisticExpr{},
	)
	if err != nil {
		return err
	}
	return nil
}

func isDatabasePopulated() bool {
	var count int64
	DB.Model(&ArticleWord{}).Count(&count)
	return count > 0
}
