package models

// MakeMigrations executes all migrations for our structs
func MakeMigrations() {
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Comment{})
}
