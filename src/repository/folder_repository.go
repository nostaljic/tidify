package repository

import (
	_ "errors"
	"fmt"
	"tidify/devlog"
	"tidify/models"

	"gorm.io/gorm"
)

type FolderRepository struct {
	DB *gorm.DB
}

func NewFolderRepository(db *gorm.DB) FolderRepository {
	return FolderRepository{
		DB: db,
	}
}

func (f FolderRepository) Migrate() error {
	devlog.Debug("[FolderRepository] - Migrate")
	return f.DB.AutoMigrate(&models.Folder{})
}

func (f FolderRepository) FindFolderListCount(email string, keyword string) (int64, error) {
	devlog.Debug("[FolderRepository] - FindFolderListCount")
	var count int64
	qry := subQueryToGetList(f.DB, email, keyword)
	err := qry.Count(&count).Error
	return count, err
}

func (f FolderRepository) FindFolderList(email string, start int, count int, keyword string) (folders []models.Folder, err error) {
	devlog.Debug("[FolderRepository] - FindFolderList")
	qry := subQueryToGetList(f.DB, email, keyword)
	err = qry.Debug().Order("updated_at DESC").Offset(start).Limit(count).Find(&folders).Error
	return folders, err

}

func subQueryToGetList(db *gorm.DB, email string, keyword string) *gorm.DB {
	qry := db.
		Model(&models.Folder{}).
		Preload("folders").
		Where("user_email=?", email)
	if len(keyword) > 0 {
		qry = qry.Where("folder_title LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}
	return qry
}

func (f FolderRepository) Create(folder *models.Folder) error {
	if err := f.DB.Create(folder).Error; err != nil {
		return err
	}
	return nil
	//TODO: RESPONSE
}

func (f FolderRepository) Delete(folderId int) error {
	folder := models.Folder{}
	if err := f.DB.Debug().
		Model(&folder).Where("folder_id = ?", folderId).Delete(&folder).Error; err != nil {
		return err
	}
	return nil
	//TODO: RESPONSE
}

func (f FolderRepository) Update(folder *models.Folder) error {
	if err := f.DB.Debug().
		Model(&folder).Where("folder_id = ?", folder.FolderID).Updates(folder).Error; err != nil {
		return err
	}
	return nil
	//TODO: RESPONSE
}

func (f FolderRepository) GetFolderByID(folderId int) *models.Folder {
	myFolder := models.Folder{}
	if err := f.DB.Debug().
		Model(models.Folder{}).
		Where("folder_id=?", folderId).First(&myFolder).Error; err != nil {
		return nil
	}
	return &myFolder
}
