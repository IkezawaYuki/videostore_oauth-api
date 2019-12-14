package access_token

import "time"

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID 		int64  `json:"user_id"`
	ClientID	int64  `json:"client_id"`
	Expire 		int64  `json:"expire"`
}

const (
	expiration = 24
)

func GetNewAccessToken() *AccessToken{
	return &AccessToken{
		AccessToken: "",
		UserID:      0,
		ClientID:    0,
		Expire:      time.Now().UTC().Add(time.Hour * expiration).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return false
}