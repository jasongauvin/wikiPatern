package export

import (
	"fmt"
	"github.com/jasongauvin/wikiPattern/models"
)

type Csv struct {
}

func (c *Csv) export(article *models.Article) *ArticleExportFile {
	fmt.Println("Export in csv")
	articleExportFile := new(ArticleExportFile)

	return articleExportFile
}
