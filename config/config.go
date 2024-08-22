package config

const (
	PORT             = 7070
	DB_FOLDER string = "./db"
	DB_PATH   string = DB_FOLDER + "/library.db"
)

func Initialize() error {
	/*
		err := InitializeDatabase()
		if err != nil {
			return err
		}
	*/
	return nil
}
