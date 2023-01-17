package repository

import (
	"book-api/db"
	model "book-api/model"
	"errors"

	"gorm.io/gorm"
)

func UserAvail(cred model.User) error {
	if err := db.DB.Where("username = ?", cred.Username).Where("password = ?", cred.Password).First(&cred).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	return nil
}
