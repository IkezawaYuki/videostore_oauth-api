package db

import (
	"github.com/IkezawaYuki/videostore_oauth-api/src/clients/cassandra"
	"github.com/IkezawaYuki/videostore_oauth-api/src/domain/access_token"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
)

func NewRepository()DbRepository{
	return &dbRepository{}
}

type DbRepository interface {
	GetByID(string)(*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetByID(id string)(*access_token.AccessToken, *errors.RestErr){
	session, err := cassandra.GetSession()
	if err != nil{
		panic(err)
	}
	defer session.Close()

	return nil, errors.NewInternalServerErr("database connection not implemented yet")
}

