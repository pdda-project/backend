package users

import "gorm.io/gorm"

type User struct {
	// PK
	Uid string `gorm:"column:uid;primaryKey"`

	// Credential
	Email        string `gorm:"column:email;unique"`
	PasswordHash string `gorm:"column:password_hash"`

	// Information
	FirstName *string `gorm:"column:first_name"`
	LastName  *string `gorm:"column:last_name"`
	Company   *string `gorm:"column:company"`
	Role      *string `gorm:"column:role"`

	// Permission
	IsAdmin  bool `gorm:"column:is_admin;default:false"`
	IsStaff  bool `gorm:"column:is_staff;default:false"`
	IsExpert bool `gorm:"column:is_expert;default:false"`

	// Metadata
	CreatedAt int            `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int            `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`

	// Relations
}
