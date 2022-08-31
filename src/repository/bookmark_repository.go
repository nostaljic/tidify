package repository

import (
	_ "errors"
	"fmt"
	"tidify/devlog"
	"tidify/models"

	"gorm.io/gorm"
)

type BookmarkRepository struct {
	DB *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) BookmarkRepository {
	return BookmarkRepository{
		DB: db,
	}
}

func (f BookmarkRepository) Migrate() error {
	devlog.Debug("[BookmarkRepository] - Migrate")
	return f.DB.AutoMigrate(&models.Bookmark{})
}

func (f BookmarkRepository) FindBookmarkListCount(email string, keyword string, folderId int) (int64, error) {
	devlog.Debug("[BookmarkRepository] - FindBookmarkListCount")
	var count int64
	qry := subQueryToGetBookmarkList(f.DB, email, keyword, folderId)
	err := qry.Count(&count).Error
	return count, err
}

func (f BookmarkRepository) FindBookmarkList(email string, start int, count int, keyword string, folderId int) (Bookmarks []models.Bookmark, err error) {
	devlog.Debug("[BookmarkRepository] - FindBookmarkList")
	qry := subQueryToGetBookmarkList(f.DB, email, keyword, folderId)
	err = qry.Debug().Order("updated_at DESC").Offset(start).Limit(count).Find(&Bookmarks).Error
	return Bookmarks, err

}

func subQueryToGetBookmarkList(db *gorm.DB, email string, keyword string, folderId int) *gorm.DB {
	qry := db.
		Model(&models.Bookmark{}).
		Preload("bookmarks").
		Where("user_email=?", email)

	if folderId != 0 {
		qry = qry.Where("folder_id = ? ", folderId)
	}
	if len(keyword) > 0 {
		qry = qry.Where("bookmark_title LIKE ? OR bookmark_url LIKE ?", fmt.Sprintf("%%%s%%", keyword), fmt.Sprintf("%%%s%%", keyword))
	}
	return qry
}

func (f BookmarkRepository) Create(Bookmark *models.Bookmark) error {
	if err := f.DB.Create(Bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (f BookmarkRepository) Delete(BookmarkId int) error {
	Bookmark := models.Bookmark{}
	if err := f.DB.Debug().
		Model(&Bookmark).Where("bookmark_id = ?", BookmarkId).Delete(&Bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (f BookmarkRepository) Update(Bookmark *models.Bookmark) error {
	if err := f.DB.Debug().
		Model(&Bookmark).Where("bookmark_id = ?", Bookmark.BookmarkID).Updates(Bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (f BookmarkRepository) GetBookmarkByID(BookmarkId int) *models.Bookmark {
	myBookmark := models.Bookmark{}
	if err := f.DB.Debug().
		Model(models.Bookmark{}).
		Where("bookmark_id=?", BookmarkId).First(&myBookmark).Error; err != nil {
		return nil
	}
	return &myBookmark
}
