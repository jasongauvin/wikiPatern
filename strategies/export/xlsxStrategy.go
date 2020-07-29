package export

import (
	"bytes"
	"github.com/jasongauvin/wikiPattern/models"
	"github.com/tealeg/xlsx"
	"strconv"
	"time"
)

// Xlsx is the csv export strategy struct
type Xlsx struct {
}

func (x *Xlsx) export(article *models.Article) (*ArticleExportFile, error) {
	var b bytes.Buffer
	var err error

	wb, err := hydrateRows(article)

	if err != nil {
		return nil, err
	}

	articleExportFile := new(ArticleExportFile)
	if err := wb.Write(&b); err != nil {
		return nil, err
	}

	// Fill in article file struct
	articleExportFile.FileBytes = b.Bytes()
	articleExportFile.MimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	articleExportFile.FileName = time.Now().UTC().Format("data-20200728173805.xlsx")

	return articleExportFile, nil
}

func createSheet() (*xlsx.Sheet, *xlsx.File, error) {
	wb := xlsx.NewFile()
	sh, err := wb.AddSheet("Articles")
	if err != nil {
		return sh, wb, err
	}
	return sh, wb, nil
}

func hydrateRows(article *models.Article) (*xlsx.File, error) {
	sh, wb, err := createSheet()

	if err != nil {
		return nil, err
	}

	row := sh.AddRow()
	row.AddCell().Value = "Id"
	row.AddCell().Value = "Title"
	row.AddCell().Value = "Content"
	row.AddCell().Value = "Created at"
	row2 := sh.AddRow()
	row2.AddCell().Value = strconv.FormatUint(article.ID, 10)
	row2.AddCell().Value = article.Title
	row2.AddCell().Value = article.Content
	row2.AddCell().Value = time.Time.String(article.CreatedAt)

	return wb, nil
}
