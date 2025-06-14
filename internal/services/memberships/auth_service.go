package memberships

import (
	"gorm.io/gorm"

	"github.com/rs/zerolog/log"

	"catalog-music/internal/models/memberships"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (r *service) SignUp(req *memberships.SignUpRequest) error {
	existingUser, err := r.userRepo.GetUser(0, req.Email, req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Error getting user")
		return err
	}

	if existingUser != nil {
		return errors.New("username or email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password")
		return err
	}

	model := &memberships.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return r.userRepo.CreateUser(model)
}
