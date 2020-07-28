package export

import (
	"fmt"
	"github.com/jasongauvin/wikiPattern/models"
	"github.com/tealeg/xlsx"
	"log"
)

type Xlsx struct {
}

func (x *Xlsx) export(article models.Article) {
	/*article, err := models.FindArticleByID(2)
	fmt.Println(article)
	if err != nil {
		log.Fatal("Cannot find article", err)
	}*/
	HydrateRows(article)
}

func CreateSheet() (*xlsx.Sheet, *xlsx.File) {
	fmt.Println("Export in csv")
	wb := xlsx.NewFile()
	sh, err := wb.AddSheet("Articles")
	if err != nil {
		log.Fatal("An error occurred while creating sheet", err)
	}
	fmt.Println(sh)
	return sh, wb
}

func HydrateRows (article models.Article){
	sh, wb := CreateSheet()
	row := sh.AddRow()
	row.AddCell().Value = "Id"
	row.AddCell().Value = "Title"
	row.AddCell().Value = "Content"
	err := wb.Save("/Users/jackmaarek/go/src/github.com/jasongauvin/wikiPattern/")
	if err != nil {
		fmt.Println("Error while saving file", err)
	}
}