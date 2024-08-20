package config

import (
	"os"

	schemas "github.com/joeCavZero/library-rest-api-in-golang/internal/schema"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/util/logkit"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDatabase() error {
	logger := logkit.New("sqlite initialization")

	var err error

	// Check if database already exists
	_, err = os.Stat(DB_PATH)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Debugf("Creating database folder %s", DB_FOLDER)

			// creating the folder for store the database
			err = os.MkdirAll(DB_FOLDER, os.ModePerm)
			if err != nil {
				logger.Error(err)
				return err
			}

			// creating the .db database file
			file, err := os.Create(DB_PATH)
			if err != nil {
				logger.Error(err)
				return err
			}
			file.Close()

			logger.Debugf("Database %s created", DB_PATH)

		}
	}

	// Open a connection to the SQLite database
	Database, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	if err != nil {
		logger.Error(err)
		return err
	}

	// AutoMigrate will create the table based on the schema provided.
	err = Database.AutoMigrate(&schemas.BookModel{})
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
