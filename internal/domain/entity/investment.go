package entity

import "time"

type Investment struct {
	ID           string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID       string         `gorm:"type:uuid;notnull" json:"user_id"`
	User         User           `gorm:"foreignKey:UserID" json:"-"`
	AccountID    *string        `gorm:"type:uuid" json:"account_id"`
	Account      *Account       `gorm:"foreignKey:AccountID" json:"-"`
	Description  string         `json:"description"`
	Type         InvestmentType `json:"investment_type"`
	Value        int            `json:"value"`
	Date         time.Time      `json:"date"`
	AnnualIncome float64        `json:"annual_income"`
	CreatedAt 	 time.Time  	`json:"created_at"`
	UpdatedAt 	 *time.Time 	`json:"updated_at"`
	DeletedAt 	 *time.Time 	`json:"deleted_at"`
}