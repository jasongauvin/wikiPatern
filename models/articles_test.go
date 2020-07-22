package models

import (
	"database/sql"
	"github.com/go-testfixtures/testfixtures/v3"
	"testing"

	//"github.com/jasongauvin/wikiPattern/models"
	//"github.com/jinzhu/gorm"
	"log"
)

func TestMain(m *testing.M) {
	log.Printf("db")
	var err error


	db, err := sql.Open("sql", "dbname=wiki")
	if err != nil {
		log.Printf("Saucisse")
	}

	_, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Files(
			"fixtures/Articles.yml",
		),
	)

	if err != nil {
		log.Printf("Pute")
	}
}