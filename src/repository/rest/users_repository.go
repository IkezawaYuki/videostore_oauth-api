package rest

import (
	"github.com/IkezawaYuki/videostore_users-api/domain/users"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	return nil, nil
}
