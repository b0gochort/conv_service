package infrastucture

import (
	"fmt"

	"github.com/b0gochort/conv_service/internal/entity"
	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	db *pg.DB
}

func NewUserRepo(db *pg.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// сюда дойдут данные уже в том формате, который можно вставлять в базу данных
func (r *UserRepo) Create(user *entity.User) error {
	_, err := r.db.Model(user).Returning("id").Insert()
	if err != nil {
		return fmt.Errorf("UserRepo - Create - DB.Exec,%w", err)
	}

	return nil
}

func (r *UserRepo) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.Model(user).Where("email = ?", email).Select()
	if err != nil {
		return nil, err
	}
	return user, nil

}
