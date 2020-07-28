package validators

import (
	"errors"
	"github.com/jasongauvin/wikiPattern/services"
)

func ValidateArticle(article *services.ArticleForm) error {
	if article.Title =="" {
		return errors.New("Required title")
	}
	if article.Content =="" {
		return errors.New("Required content")
	}
	return nil
}