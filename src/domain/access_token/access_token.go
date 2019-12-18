package access_token

import (
	"github.com/IkezawaYuki/videostore_users-api/utils/errors"
	"strings"
	"time"
)

const (
	expiration = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID 		int64  `json:"user_id"`
	ClientID	int64  `json:"client_id"`
	Expires 		int64  `json:"expire"`
}

func (at *AccessToken) Validate() *errors.RestErr{
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == ""{
		return errors.NewBadRequestErr("invalid access token id")
	}
	if at.UserID <= 0{
		return errors.NewBadRequestErr("invalid user id")
	}
	if at.ClientID <= 0{
		return errors.NewBadRequestErr("invalid client id")
	}
	if at.Expires <= 0{
		return errors.NewBadRequestErr("invalid expiration id")
	}
	return nil
}

func GetNewAccessToken() *AccessToken{
	return &AccessToken{
		Expires:      time.Now().UTC().Add(time.Hour * expiration).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}