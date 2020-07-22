package models

func MakeMigrations() {
	db.AutoMigrate(&Post{})
}