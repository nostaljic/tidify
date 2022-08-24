package models

type Folder struct {
	BaseModel
	UserID      uint   `gorm:"not null" json:"user_id"`
	FolderID    uint   `gorm:"primary_key" json:"folder_id"`
	FolderTitle string `json:"folder_title" gorm:"not null"`
	FolderColor string `json:"folder_color"`
}
