package memberships

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Email     string `gorm:"not null;unique"`
		Username  string `gorm:"not null;unique"`
		Password  string `gorm:"not null"`
		CreatedBy string `gorm:"not null"`
		UpdatedBy string `gorm:"not null"`
	}
)

type (
	SignUpRequest struct {
		Email    string `json:"email" binding:"required,email,min=3,max=32"`
		Username string `json:"username" binding:"required,min=3,max=32"`
		Password string `json:"password" binding:"required,min=8"`
	}
)

func (u *User) TableName() string {
	return "users"
}
