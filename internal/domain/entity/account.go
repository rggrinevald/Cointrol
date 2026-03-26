package entity

import "time"

type Account struct {
	ID        string      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string      `gorm:"type:uuid;notnull" json:"user_id"`
	User      User        `gorm:"foreignKey:UserID" json:"-"`
	Name      string      `json:"name"`
	Bank      BankAccount `json:"bank"`
	Type      AccountType `json:"type"`
	CardLimit *int        `json:"card_limit"`
	MatureDay *int        `json:"mature_day"`
	CloseDay  *int        `json:"close_day"`
	Balance   *int        `json:"balance"`
    CreatedAt time.Time   `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
}