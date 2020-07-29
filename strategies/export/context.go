package export

import (
	"github.com/jasongauvin/wikiPattern/models"
	"github.com/jasongauvin/wikiPattern/services"
	"log"
)

// Context is the structure for an export, based on the exportInterface
type Context struct {
	exportInterface exportInterface
}

// ArticleExportFile is the structure used to generate a response body for an article export file
type ArticleExportFile struct {
	FileBytes []byte
	MimeType  string
	FileName  string
}

// NewContext is the constructor of a Context for exporting article
func NewContext(e exportInterface) *Context {
	return &Context{
		exportInterface: e,
	}
}

// Export function is the context function that generate a file from an article id
func (c *Context) Export(param string) *ArticleExportFile {
	var articleExportFile *ArticleExportFile
	var err error
	var article *models.Article

	article, err = services.LoadArticleById(param)

	if err != nil {
		log.Fatal("An error occurred while fetching data: ", err)
	}

	articleExportFile, err = c.exportInterface.export(article)
	if err != nil {
		log.Fatal("An error occured while exporting article:", err)
	}

	return articleExportFile
}
