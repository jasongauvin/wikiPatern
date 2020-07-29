package tests

import (
	"testing"
	"time"
	"strconv"
	"github.com/caarlos0/env/v6"
	"log"
	"github.com/jasongauvin/wikiPattern/strategies/export"
	"github.com/jasongauvin/wikiPattern/models"
)

type config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     int    `env:"DB_PORT" envDefault:"3306"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
}

func prepareDB() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	
	// Database initialization
	models.InitializeDb(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, cfg.DbPort)
	models.MakeMigrations()
}

func createArticle() (models.Article, error) {
	article := models.Article {
		Title: "Test article 1",
		Content: "This is the content of the test article 1.",
		CreatedAt: time.Now(),
	}
	err := models.CreateArticle(&article)

	if err != nil {
		return article, err
	}

	return article, nil
}

func TestCsvExport(t *testing.T) {
	var article models.Article
	var err error
	
	prepareDB()

	// Create test article
	article, err = createArticle()

	if err != nil {
		t.Errorf("Couldn't create article for the test")
	}

	csv := &export.Csv{}
	exportContext := export.NewContext(csv)

	_ = exportContext.Export(strconv.FormatUint(article.ID, 10))
}

func TestXlsxExport(t *testing.T) {
	var article models.Article
	var err error
	
	prepareDB()

	// Create test article
	article, err = createArticle()

	if err != nil {
		t.Errorf("Couldn't create article for the test")
	}

	xlsx := &export.Xlsx{}
	exportContext := export.NewContext(xlsx)

	_ = exportContext.Export(strconv.FormatUint(article.ID, 10))
}

