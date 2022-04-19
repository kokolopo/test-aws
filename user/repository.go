package user

import "gorm.io/gorm"

type UserRepository interface {
	Save(user User) (User, error)
	FetchAll() ([]User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FetchAll() ([]User, error) {
	var user []User

	err := r.db.Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
