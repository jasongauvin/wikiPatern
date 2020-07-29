package export

import (
	"fmt"
	"github.com/jasongauvin/wikiPattern/models"
	"encoding/csv"
	"bytes"
	"time"
	"strconv"
)

type Csv struct {
}

func (c *Csv) export(article *models.Article) *ArticleExportFile {
	fmt.Println("Export in csv")
	articleExportFile := new(ArticleExportFile)

	b := createCsv(article)

	articleExportFile.FileBytes = b.Bytes()
	articleExportFile.MimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	articleExportFile.FileName = time.Now().UTC().Format("data-20200728173805.csv")

	return articleExportFile
}

func createCsv(article *models.Article) (bytes.Buffer) {
	var b bytes.Buffer
	writer := csv.NewWriter(&b)
	data := [][]string{{"Id", "Title", "Content", "CreatedAt"}, {strconv.FormatUint(article.ID, 10), article.Title, article.Content, time.Time.String(article.CreatedAt)}}
	
	writer.WriteAll(data)
	
	writer.Flush()

	return b
}
