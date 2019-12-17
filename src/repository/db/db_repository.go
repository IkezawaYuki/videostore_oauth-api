package db

import (
	"github.com/IkezawaYuki/videostore_oauth-api/src/clients/cassandra"
	"github.com/IkezawaYuki/videostore_oauth-api/src/domain/access_token"
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"github.com/gocql/gocql"
)

const(
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
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

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
		); err != nil{
			if err == gocql.ErrNotFound{
				return nil, errors.NewNotFoundErr("no access token found with given id")
			}
		return nil, errors.NewInternalServerErr(err.Error())
	}
	return &result, nil
}

