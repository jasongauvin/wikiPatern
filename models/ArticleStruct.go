package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	ID   uint64 `gorm:"primary_key"`
	Title string `gorm:"size:255"`
	Content string `gorm:"size:2000"`
	CreatedAt *time.Time `gorm: not null`
	UpdatedAt *time.Time
}

func FindArticleByID(uid uint64) (Article, error) {
	var err error
	var article Article
	err = db.Debug().First(&article, uid).Error
	if err != nil {
		return Article{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return Article{}, errors.New("Article Not Found")
	}
	return article, nil
}

func FindArticles() ([]Article, error) {
	var err error
	var articles []Article
	err = db.Debug().Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func DeleteArticleByID(id uint64) error {
	var err error
	var article Article

	err = db.Debug().First(&article, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Article Not Found")
	}
	err = db.Debug().Delete(&article, id).Error
	if err != nil {
		return err
	}

	return nil
}

func EditArticleByID(article *Article, id uint64) error {
	var err error
	var old Article
	err = db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Article Not Found")
	}
	article.ID = id
	err = db.Debug().Save(&article).Error
	if err != nil {
		return errors.New("Could'nt update article")
	}
	return nil
}

// CreatedArticle creates an article row in database
func CreateArticle(article *Article) error {
	var err error
	err = db.Debug().Create(article).Error

	if err != nil {
		return err
	}
	return nil
}