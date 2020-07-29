package export

import "github.com/jasongauvin/wikiPattern/models"

type exportInterface interface {
	export(article *models.Article) (*ArticleExportFile, error)
}
