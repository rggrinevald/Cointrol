package entity

import "time"

type User struct {
	ID        string     `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `json:"-"`
	Role      UserRole   `json:"role"`
	Plan      UserPlan   `json:"plan"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}