package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	ID       uint64     `gorm:"primary_key"`
	Name     string     `gorm:"size:255;unique;not null"`
	Articles []*Article `gorm:"many2many:tags_articles;"`
}

// FindTagByID allows you to find a specific tag using its id
func FindTagByID(uid uint64) (Tag, error) {
	var err error
	var tag Tag
	err = db.Debug().First(&tag, uid).Error
	if err != nil {
		return Tag{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return Tag{}, errors.New("Tag Not Found")
	}
	return tag, nil
}

// FindTags returns you a list of tags
func FindTags() ([]Tag, error) {
	var err error
	var tags []Tag
	err = db.Debug().Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// DeleteTagByID allows you to remove an tag from the db using its id
func DeleteTagByID(id uint64) error {
	var err error
	var tag Tag

	err = db.Debug().First(&tag, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Tag Not Found")
	}
	err = db.Debug().Delete(&tag, id).Error
	if err != nil {
		return err
	}

	return nil
}

// EditTagByID allow you to modify an tag using its id
func EditTagByID(tag *Tag, id uint64) error {
	var err error
	var old Tag
	err = db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Tag Not Found")
	}
	tag.ID = id

	err = db.Debug().Save(&tag).Error
	if err != nil {
		return errors.New("Could'nt update tag")
	}
	return nil
}

// CreateTag creates an tag row in database
func CreateTag(tag *Tag) error {
	var err error
	err = db.Debug().Create(tag).Error

	if err != nil {
		return err
	}
	return nil
}

func FindTagByName(name string) (Tag, error) {
	var err error
	var tag Tag
	err = db.Debug().Where("name = ?", name).First(&tag).Error
	if err != nil {
		return Tag{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return Tag{}, errors.New("Tag Not Found")
	}
	return tag, nil
}
