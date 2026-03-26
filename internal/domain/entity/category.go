package entity

import "time"

type Category struct {
	ID        string       `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    *string      `gorm:"type:uuid" json:"user_id"`
	User      *User        `gorm:"foreignKey:UserID" json:"-"`
	Name      string       `json:"name"`
	Type      CategoryType `json:"type"`
	Icon      string       `json:"icon"`
	Color     string       `json:"color"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt *time.Time   `json:"updated_at"`
	DeletedAt *time.Time   `json:"deleted_at"`
}