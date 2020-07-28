package validators
import (
"errors"
"github.com/jasongauvin/wikiPattern/services"
)

func ValidateComment(comment *services.CommentForm) error {
	if comment.Content =="" {
		return errors.New("Required content")
	}
	return nil
}