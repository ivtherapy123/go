package UserService

import (
	"gorm.io/gorm"
	"strconv"
)

type MessageRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var tasks []User
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {

	var existingUser User
	r.db.First(&existingUser, id)

	existingUser.Email = user.Email
	existingUser.Password = user.Password

	r.db.Save(&existingUser)
	return existingUser, nil
}

func (r *userRepository) DeleteUserByID(id string) error {
	var user User
	uintId, err := strconv.ParseUint(id, 10, 32) // Преобразуем в uint
	if err != nil {
		return err
	}
	r.db.First(&user, uintId)
	r.db.Delete(&user)
	return nil
}
