package export

import (
	"fmt"
	"github.com/jasongauvin/wikiPattern/models"
	"github.com/jasongauvin/wikiPattern/services"
	"log"
)

type ExportContext struct {
	exportInterface exportInterface
}

type ArticleExportFile struct {
	FileBytes []byte
	MimeType  string
	FileName  string
}

func InitExportContext(e exportInterface) *ExportContext {
	return &ExportContext{
		exportInterface: e,
	}
}

func (c *ExportContext) SetExportInterface(e exportInterface) {
	c.exportInterface = e
}

func (c *ExportContext) Export(param string) *ArticleExportFile {
	var articleExportFile *ArticleExportFile
	var err error
	var article *models.Article
	article, err = services.LoadArticleById(param)
	fmt.Println(article)
	if err != nil {
		log.Fatal("An error occurred while fetching data: ", err)
	}
	articleExportFile = c.exportInterface.export(article)
	if articleExportFile == nil {
		log.Fatal("Export file was not generated:", err)
	}
	return articleExportFile
}
