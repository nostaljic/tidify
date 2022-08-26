package repository

import (
	_ "errors"
	"tidify/devlog"
	"tidify/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}
func (f UserRepository) Migrate() error {
	devlog.Debug("[UserRepository] - Migrate")
	return f.DB.AutoMigrate(&models.User{})
}

func (f UserRepository) IsUserExist(user *models.User) (bool, *DBError) {
	userResult := &models.User{}

	if err := f.DB.Where("user_email=? AND sns_type=?", user.UserEmail, user.SnsType).First(userResult).Error; err != nil {
		return false, CreateDBError(DB_INFRA_ERROR, "Can't find Database")
	}
	if len(userResult.UserEmail) != 0 {
		devlog.Debug("[IsUserExist] Can Register", userResult)
		return false, nil
	}
	return true, nil

}
func (f UserRepository) Create(user *models.User) (bool, *DBError) {

	if err := f.DB.Create(user).Error; err != nil {
		return false, CreateDBError(DB_INFRA_ERROR, "Can't find Database")
	}
	return true, nil
}
func checkEmailHash(hashVal, userPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))
	if err != nil {
		return false
	} else {
		return true
	}
}
