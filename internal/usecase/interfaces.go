package usecase

import "github.com/b0gochort/conv_service/internal/entity"

type (
	User interface {
		SignUp(user entity.User) error
		LogIn(email, password string) (entity.User, error)
	}
	UserRepo interface {
		Create(user *entity.User) error
		GetUserByEmail(email string) (*entity.User, error)
	}
)
