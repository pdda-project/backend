// users/services.go

package users

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	httpErr "github.com/pdda-project/backend/libs/errors"
	"github.com/pdda-project/backend/libs/validation"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService
type UserService struct {
	db *gorm.DB
}

// UserService Constructor
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// Register user service
func (s *UserService) register(data *ReqRegister) (*ResRegister, *httpErr.HTTPError) {

	// ========================= Count user
	if count, err := s._countUser(data.Email); err != nil {
		// Handler database error
		return nil, httpErr.NewHTTPError(500, []string{"database error on count user"}, err)
	} else if count != 0 {
		// Handle conflict email
		return nil, httpErr.NewHTTPError(409, []string{"conflict error on count user", "email already exist"}, nil)
	}

	// ========================= Validate data
	if _, err := validation.ValidateStruct(data); err != nil {
		if valErr, ok := err.(validator.ValidationErrors); ok {
			// Handle validation error
			validationErrors := []string{"validation error on validate register user", valErr.Error()}
			for _, f := range valErr {
				validationErrors = append(validationErrors, fmt.Sprintf("invalid %s", f.Tag()))
			}
			return nil, httpErr.NewHTTPError(400, validationErrors, valErr)

		} else {
			// Handle other error from validator
			return nil, httpErr.NewHTTPError(500, []string{"validator error on validate register user", err.Error()}, err)

		}
	}
	// ========================= Create user instance
	// Generate uuid
	newUuid, err := uuid.NewV7()
	if err != nil {
		return nil, httpErr.NewHTTPError(500, []string{"uuid error on register user", err.Error()}, err)
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return nil, httpErr.NewHTTPError(500, []string{"bcrypt error on register user", err.Error()}, err)

	}

	user := &User{
		Uid:          newUuid.String(),
		Email:        data.Email,
		PasswordHash: string(passwordHash),
	}
	// ========================= Insert to database

	if err := s.db.Create(user).Error; err != nil {
		return nil, httpErr.NewHTTPError(500, []string{"database error", err.Error()}, err)

	}

	return &ResRegister{
		Uid: user.Uid,
	}, nil

}

// count user by email
func (s *UserService) _countUser(email string) (int64, error) {
	var count int64
	err := s.db.Model(&User{}).Where("email = ?", email).Count(&count).Error

	return count, err
}
