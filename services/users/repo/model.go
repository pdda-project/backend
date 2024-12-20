package repo

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Uid           string  `gorm:"column:uid;primaryKey"`
	Email         string  `gorm:"column:email"`
	PasswordHash  string  `gorm:"column:password_hash"`
	DisplayName   string  `gorm:"column:display_name"`
	IsAdmin       bool    `gorm:"column:is_admin;default:false"`
	IsStaff       bool    `gorm:"column:is_staff;default:false"`
	IsExpert      bool    `gorm:"column:is_expert;default:false"`
	RoleAtCompany *string `gorm:"column:role_at_company"`
	Company       *string `gorm:"column:company"`
	CompanyEmail  *string `gorm:"column:company_email"`
	CreatedAt     string  `gorm:"column:created_at"`
	PhotoUrl      string  `gorm:"column:photo_url"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Create UID
	uid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.Uid = uid.String()

	// Hash Password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), 10)
	if err != nil {
		return err
	}
	u.PasswordHash = string(passwordHash)

	return nil
}
