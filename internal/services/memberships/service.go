package memberships

import (
	"catalog-music/internal/configs"
	"catalog-music/internal/models/memberships"
)

type membershipRepository interface {
	CreateUser(user *memberships.User) error
	GetUser(id uint, email, username string) (*memberships.User, error)
}

type service struct {
	cfg      *configs.Config
	userRepo membershipRepository
}

func NewService(cfg *configs.Config, userRepo membershipRepository) *service {
	return &service{cfg: cfg, userRepo: userRepo}
}
