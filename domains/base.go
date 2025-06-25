package domains

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel contains common fields for all domain models
type BaseModel struct {
	ID        string     `gorm:"primary_key;type:varchar(36)" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// BeforeCreate hook to generate UUID
func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return
}

// User represents a user in the system
type User struct {
	BaseModel
	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Email       string `gorm:"column:email;type:varchar(100);not null;uniqueIndex" json:"email"`
	Username    string `gorm:"column:username;type:varchar(100);uniqueIndex" json:"username,omitempty"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(20)" json:"phone_number,omitempty"`
	IsActive    bool   `gorm:"column:is_active;default:true" json:"is_active"`
}

// TableName sets the table name for User model
func (User) TableName() string {
	return "users"
}
