package config

import (
	"gorm.io/gorm"
)

const (
	PORT             = 7070
	DB_FOLDER string = "./db"
	DB_PATH   string = DB_FOLDER + "/library.db"
)

var (
	Database *gorm.DB
)

func Initialize() error {

	err := InitializeDatabase()
	if err != nil {
		return err
	}

	return nil
}
