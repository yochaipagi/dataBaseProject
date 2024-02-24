package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

const (
	dsnKey = "DSN"
)

var DB *gorm.DB

func SetupDB() error {
	var err error
	if DB, err = gorm.Open(postgres.Open(os.Getenv(dsnKey)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}); err != nil {
		return err
	}

	if err = DB.AutoMigrate(&Article{}, &ArticleLine{}, &ArticleWord{}, &WordGroup{}, &Word{}, &LinguisticExpr{}); err != nil {
		return err
	}

	// don't populate if there's already data
	if res := DB.First(&ArticleWord{}); res.Error == nil {
		return nil
	}

	if err = populateDB(); err != nil {
		return err
	}

	return nil
}
