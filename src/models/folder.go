package models

type Folder struct {
	BaseModel
	UserEmail   string `json:"user_email" gorm:"index;not null;check:user_email <> ''"`
	FolderID    uint   `json:"folder_id" gorm:"primary_key"`
	FolderTitle string `json:"folder_title" gorm:"not null;check:folder_title <> ''"`
	FolderColor string `json:"folder_color" gorm:"not null;check:folder_color <> ''"`
}
