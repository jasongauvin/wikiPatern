package export

import (
	"bytes"
	"fmt"
	"github.com/jasongauvin/wikiPattern/models"
	"github.com/tealeg/xlsx"
	"log"
	"strconv"
	"time"
)

type Xlsx struct {
}

func (x *Xlsx) export(article *models.Article) *ArticleExportFile {
	var b bytes.Buffer
	wb := hydrateRows(article)
	articleExportFile := new(ArticleExportFile)
	if err := wb.Write(&b); err != nil {
		log.Fatal("An error occurred while creating sheet", err)
	}
	articleExportFile.FileBytes = b.Bytes()
	articleExportFile.MimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	articleExportFile.FileName = time.Now().UTC().Format("data-20200728173805.xlsx")

	return articleExportFile
}

func createSheet() (*xlsx.Sheet, *xlsx.File) {
	fmt.Println("Export in csv")
	wb := xlsx.NewFile()
	sh, err := wb.AddSheet("Articles")
	if err != nil {
		log.Fatal("An error occurred while creating sheet", err)
	}
	return sh, wb
}

func hydrateRows(article *models.Article) *xlsx.File {
	sh, wb := createSheet()
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

	return wb
}
