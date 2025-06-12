package memberships

import "catalog-music/internal/models/memberships"

func (r *repository) CreateUser(user *memberships.User) error {
	return r.db.Create(user).Error
}

func (r *repository) GetUser(id uint, email, username string) (memberships.User, error) {
	var user memberships.User
	if err := r.db.Where("id = ? OR email = ? OR username = ?", id, email, username).First(&user).Error; err != nil {
		return memberships.User{}, err
	}
	return user, nil
}
