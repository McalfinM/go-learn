package account

import (
	"learn/api/databases/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.User, error)
	FindById(Id int) (model.User, error)
	Create(user model.User) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) FindById(ID int) (model.User, error) {
	var user model.User
	if err := r.db.First(&user, ID).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *repository) Create(user model.User) (model.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
