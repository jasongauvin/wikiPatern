package models

func MakeMigrations() {
	db.AutoMigrate(&Article{})
}