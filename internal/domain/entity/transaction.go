package entity

import "time"

type Transaction struct {
	ID            string        `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID        string        `gorm:"type:uuid;notnull" json:"user_id"`
	User          User          `gorm:"foreignKey:UserID" json:"-"`
	AccountID     string        `gorm:"type:uuid;notnull" json:"account_id"`
	Account       Account       `gorm:"foreignKey:AccountID" json:"-"`
	CategoryID    string        `gorm:"type:uuid;notnull" json:"category_id"`
    Category      Category      `gorm:"foreignKey:CategoryID" json:"-"`
	Value         int           `json:"value"`
	Description   string        `json:"description"`
	Date          time.Time     `json:"date"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	CreatedAt 	  time.Time  	`json:"created_at"`
	UpdatedAt	  *time.Time    `json:"updated_at"`
	DeletedAt 	  *time.Time    `json:"deleted_at"`
}