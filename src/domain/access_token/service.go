package access_token

import (
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"strings"
)

type Repository interface {
	GetByID(string)(*AccessToken, *errors.RestErr)
}

type Service interface {
	GetByID(string)(*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository)(Service){
	return &service{
		repository: repo,
	}
}

func (s *service)GetByID(accessTokenID string)(*AccessToken, *errors.RestErr){
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0{
		return nil, errors.NewBadRequestErr("invalid access token id")
	}
	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil{
		return nil, err
	}
	return accessToken, nil
}