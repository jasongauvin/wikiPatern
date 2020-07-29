package export

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/jasongauvin/wikiPattern/models"
	"strconv"
	"time"
)

// Csv is the csv export strategy struct
type Csv struct {
}

func (c *Csv) export(article *models.Article) (*ArticleExportFile, error) {
	fmt.Println("Export in csv")
	articleExportFile := new(ArticleExportFile)
	
	var err error
	b, err := createCsv(article)

	if err != nil {
		return nil, err
	}

	// Fill in article file struct
	articleExportFile.FileBytes = b.Bytes()
	articleExportFile.MimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	articleExportFile.FileName = time.Now().UTC().Format("data-20200728173805.csv")

	return articleExportFile, nil
}

func createCsv(article *models.Article) (bytes.Buffer, error) {
	var b bytes.Buffer
	writer := csv.NewWriter(&b)
	data := [][]string{{"Id", "Title", "Content", "CreatedAt"}, {strconv.FormatUint(article.ID, 10), article.Title, article.Content, time.Time.String(article.CreatedAt)}}
	
	var err error
	err = writer.WriteAll(data)

	if err != nil {
		return b, err
	}

	return b, nil
}
