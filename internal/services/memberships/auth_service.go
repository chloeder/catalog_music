package memberships

import (
	"time"

	"gorm.io/gorm"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"

	"catalog-music/internal/models/memberships"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(req *memberships.SignUpRequest) error {
	existingUser, err := s.userRepo.GetUser(0, req.Email, req.Username)
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

	return s.userRepo.CreateUser(model)
}

func (s *service) SignIn(req *memberships.SignInRequest) (string, error) {
	existingUser, err := s.userRepo.GetUser(0, "", req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Error getting user")
		return "", err
	}

	if existingUser == nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte(s.cfg.Service.SecretJWT))
	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}
