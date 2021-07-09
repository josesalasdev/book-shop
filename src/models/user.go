//models.user

package models

import (
	"github.com/josesalasdev/golang_api_template/src/utils"
)

// User model.
type User struct {
	ModelBase
	Email    string `json:"user"`
	Password string `json:"password"`
}

// UserInput for userModel.
type UserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// EncryptPassword method.
func (u *User) EncryptPassword() {
	passwordEncrypt := utils.Encrypt(u.Password)
	u.Password = passwordEncrypt
}

// ValidatePassword method.
func (u *User) ValidatePassword(password string) bool {
	if err := utils.CompareHash(u.Password, password); err != nil {
		return false
	}
	return true
}
