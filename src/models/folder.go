package models

type Folder struct {
	BaseModel
	UserEmail   string `gorm:"index;not null" json:"user_email"`
	FolderID    uint   `gorm:"primary_key" json:"folder_id"`
	FolderTitle string `json:"folder_title" gorm:"not null"`
	FolderColor string `json:"folder_color" gorm:"not null"`
}
