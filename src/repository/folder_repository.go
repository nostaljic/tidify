package repository

import (
	_ "errors"
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
func (f FolderRepository) FindFolderList() (folders []models.Folder, err error) {
	devlog.Debug("[FolderRepository] - FindFolderList")
	err = f.DB.Find(&folders).Error
	devlog.Debug("[FolderRepository] - err", err)
	devlog.Debug("[FolderRepository] - folders", folders)
	return folders, err

}
func (f FolderRepository) Create(folder *models.Folder) {

	if err := f.DB.Create(folder).Error; err != nil {
		devlog.Fatal("DB Error : Folder", err)
	}
}
