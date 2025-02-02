package model

type Task struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title" gorm:"not null"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	User      User   `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	UserID    uint   `json:"user_id"`
}