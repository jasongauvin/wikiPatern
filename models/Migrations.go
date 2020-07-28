package models

// MakeMigrations executes all migrations for our structs
func MakeMigrations() {
	db.AutoMigrate(&Article{}, &Comment{}, &User{}, &UserSession{}, &Tag{})
}
