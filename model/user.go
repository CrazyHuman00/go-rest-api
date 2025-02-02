package model

import "time"

// jsonに変換した時に返す値を定義、自動的に小文字に変換される

// User ユーザー情報
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"` // gormでの設定、主キー
	Email	  string    `json:"email" gorm:"unique"`   // gormでの設定、重複を許さない
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse ユーザー情報のレスポンス
type UserResponse struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Email	  string    `json:"email" gorm:"unique"`
}