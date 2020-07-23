package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

// Comment is user post in the wiki
type Comment struct {
	ID        uint64 `gorm:"primary_key"`
	Text      string `gorm:"size:500"`
	ArticleId uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FindCommentByID allows you to find a specific comment using its id
func FindCommentByID(uid uint64) (Comment, error) {
	var err error
	var comment Comment
	err = db.Debug().First(&comment, uid).Error
	if err != nil {
		return Comment{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return Comment{}, errors.New("Comment Not Found")
	}
	return comment, nil
}

// FindComments returns you a list of comments
func FindComments() ([]Comment, error) {
	var err error
	var comments []Comment
	err = db.Debug().Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// DeleteCommentByID allows you to remove an comment from the db using its id
func DeleteCommentByID(id uint64) error {
	var err error
	var comment Comment

	err = db.Debug().First(&comment, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Comment Not Found")
	}
	err = db.Debug().Delete(&comment, id).Error
	if err != nil {
		return err
	}

	return nil
}

// EditCommentByID allow you to modify an comment using its id
func EditCommentByID(comment *Comment, id uint64) error {
	var err error
	var old Comment
	err = db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Comment Not Found")
	}
	comment.ID = id
	comment.UpdatedAt = time.Now()

	err = db.Debug().Save(&comment).Error
	if err != nil {
		return errors.New("Could'nt update comment")
	}
	return nil
}

// CreateComment creates an comment row in database
func CreateComment(comment *Comment) error {
	var err error
	comment.CreatedAt = time.Now()
	err = db.Debug().Create(comment).Error

	if err != nil {
		return err
	}
	return nil
}

/*func FindCommentByArticle(article *Article) (*[]Comment, error) {
	var err error
	var comment Comment
	var resultCount int
	dbComments := db.Model(&article).Related(&comment)
	fmt.Println(dbComments.Count(&resultCount))
	return dbComments, err
}*/
