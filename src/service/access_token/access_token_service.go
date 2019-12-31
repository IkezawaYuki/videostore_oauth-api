package access_token

import (
	"github.com/IkezawaYuki/videostore_oauth-api/src/domain/access_token"
	"github.com/IkezawaYuki/videostore_oauth-api/src/repository/db"
	"github.com/IkezawaYuki/videostore_oauth-api/src/repository/rest"
	"github.com/IkezawaYuki/videostore_oauth-api/src/utils/errors"
	"strings"
)

//type Repository interface {
//	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
//	Create(access_token.AccessToken) *errors.RestErr
//	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
//}

type Service interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetByID(accessTokenID string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestErr("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := s.restUsersRepo.LoginUser(request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	at := access_token.GetNewAccessToken(user.ID)
	at.Generate()

	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
