package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// dsnKey is the name of the environment variable that stores the database connection string.
const dsnKey = "DSN"

// DB is a global variable that will hold the database connection instance, accessible throughout the application.
var DB *gorm.DB

// SetupDB initializes the database connection, runs migrations, and populates the database if it's not already populated.
func SetupDB() error {
	// Retrieve the database connection string (DSN) from the environment variable.
	dsn := os.Getenv(dsnKey)
	// If the DSN is empty, return an error indicating that the environment variable is not set.
	if dsn == "" {
		return fmt.Errorf("environment variable %s is not set", dsnKey)
	}

	// Open the database connection using GORM with the PostgreSQL driver and configure logging to be silent.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	// If there's an error connecting to the database, return an error indicating the failure.
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Assign the connected database to the global DB variable.
	DB = db

	// Run database migrations to ensure the schema is up to date.
	err = migrateDB()
	// If there's an error during migration, return an error indicating the failure.
	if err != nil {
		return fmt.Errorf("failed to migrate the database: %v", err)
	}

	//Check if the database is already populated. If it is, skip the population step.
	if isDatabasePopulated() {
		return nil
	}

	// Populate the database with initial data if it's not already populated.
	err = populateDB()
	// If there's an error during population, return an error indicating the failure.
	if err != nil {
		return fmt.Errorf("failed to populate the database: %v", err)
	}

	// If everything succeeds, return nil indicating no errors.
	return nil
}

// migrateDB runs the GORM AutoMigrate function to ensure the database schema matches the Go struct definitions.
func migrateDB() error {
	// Automatically migrate the schema, creating or altering tables to match the struct definitions.
	err := DB.AutoMigrate(
		&Article{},
		&ArticleLine{},
		&ArticleWord{},
		&WordGroup{},
		&Word{},
		&LinguisticExpr{},
	)
	// Return any errors encountered during migration.
	return err
}

// isDatabasePopulated checks if the database is already populated by counting entries in the ArticleWord table.
func isDatabasePopulated() bool {
	var count int64
	// Count the number of entries in the ArticleWord table.
	DB.Model(&ArticleWord{}).Count(&count)
	// Return true if the count is greater than 0, indicating the database is populated.
	return count > 0
}
