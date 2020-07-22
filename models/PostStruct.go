package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Post struct {
	ID   uint64 `gorm:"primary_key"`
	Name string `gorm:"size:255"`
}

func FindPostByID(uid uint64) (Post, error) {
	var err error
	var post Post
	err = db.Debug().First(&post, uid).Error
	if err != nil {
		return Post{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return Post{}, errors.New("Post Not Found")
	}
	return post, nil
}

func FindPosts() ([]Post, error) {
	var err error
	var posts []Post
	err = db.Debug().Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func DeletePostByID(id uint64) error {
	var err error
	var post Post

	err = db.Debug().First(&post, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Post Not Found")
	}
	err = db.Debug().Delete(&post, id).Error
	if err != nil {
		return err
	}

	return nil
}

func EditPostByID(post *Post, id uint64) error {
	var err error
	var old Post
	err = db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Post Not Found")
	}
	post.ID = id
	err = db.Debug().Save(&post).Error
	if err != nil {
		return errors.New("Could'nt update post")
	}
	return nil
}

func CreatePost(post *Post) error {
	var err error
	err = db.Debug().Create(post).Error

	if err != nil {
		return err
	}
	return nil
}