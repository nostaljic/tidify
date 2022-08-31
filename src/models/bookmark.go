package models

type Bookmark struct {
	BaseModel
	UserEmail     string `json:"user_email" gorm:"index;check:user_email <> ''"`
	FolderID      uint   `json:"folder_id" gorm:"index;not null"`
	BookmarkID    uint   `json:"bookmark_id" gorm:"primary_key"`
	BookmarkUrl   string `json:"bookmark_url" gorm:"not null;check:bookmark_url <> ''" `
	BookmarkTitle string `json:"bookmark_title" gorm:"not null;check:bookmark_title <> ''"`
}
