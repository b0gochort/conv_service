package usecase

import (
	"time"

	"github.com/b0gochort/conv_service/internal/entity"
	"github.com/b0gochort/conv_service/internal/transport/request"
	"github.com/b0gochort/conv_service/internal/utils/crypto"
	"github.com/b0gochort/conv_service/internal/utils/jwt"
)

type UsersUseCase struct {
	repo          UserRepo
	cryptoService crypto.CryptoService
	jwtService    jwt.JWTService
}

func NewUsersUC(repo UserRepo, cryptoSvc crypto.CryptoService, jwtSvc jwt.JWTService) *UsersUseCase {
	return &UsersUseCase{
		repo:          repo,
		cryptoService: cryptoSvc,
		jwtService:    jwtSvc,
	}
}

func (uc *UsersUseCase) SignUp(request *request.SignUpReq) error {

	_, err := uc.repo.GetUserByEmail(request.Email)

	if err != nil {
		return err
	}

	passwordHash, err := uc.cryptoService.CreatePasswordHash(request.Password)
	if err != nil {
		return err
	}

	err = uc.repo.Create(&entity.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  passwordHash,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}
func (uc *UsersUseCase) LogIn(request *request.LogInReq) (accessToken string, err error) {
	user, err := uc.repo.GetUserByEmail(request.Email)
	if err != nil {
		return
	}

	if !uc.cryptoService.ValidatePassword(user.Password, request.Password) {
		return
	}

	accessToken, err = uc.jwtService.GenerateToken(user.ID)
	return

}
