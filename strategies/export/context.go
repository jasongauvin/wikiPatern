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

func InitExportContext(e exportInterface) *ExportContext {
	return &ExportContext{
		exportInterface: e,
	}
}

func (c *ExportContext) SetExportInterface(e exportInterface) {
	c.exportInterface = e
}

func (c *ExportContext) Export(param string) {
	var articleId uint64
	var err error
	var article models.Article
	articleId = services.ConvertStringToInt(param)
	article, err = models.FindArticleByID(articleId)
	fmt.Println(article)
	if err != nil {
		log.Fatal("An error occurred while fetching data: ", err)
	}
	c.exportInterface.export(article)
}
