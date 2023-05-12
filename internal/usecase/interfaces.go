package usecase

import (
	"github.com/b0gochort/conv_service/internal/entity"
	"github.com/b0gochort/conv_service/internal/transport/request"
)

type (
	UserUseCase interface {
		SignUp(request *request.SignUpReq) error
		LogIn(request *request.LogInReq) (string, error)
	}
	UserRepo interface {
		Create(user *entity.User) error
		GetUserByEmail(email string) (*entity.User, error)
	}
)
