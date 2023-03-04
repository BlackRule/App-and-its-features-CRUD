package user

import (
	"github.com/jinzhu/gorm"
)

// Repo interface
type Repo interface {
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo will instantiate User Repository
func NewUserRepo(db *gorm.DB) Repo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetByID(id uint) (*User, error) {
	var user User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetByEmail(email string) (*User, error) {
	var user User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) Create(user *User) error {
	return u.db.Create(user).Error
}

func (u *userRepo) Update(user *User) error {
	return u.db.Save(user).Error
}
