package access_token

import "time"

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID 		int64  `json:"user_id"`
	ClientID	int64  `json:"client_id"`
	Expires 		int64  `json:"expire"`
}

const (
	expiration = 24
)

func GetNewAccessToken() *AccessToken{
	return &AccessToken{
		AccessToken: "",
		UserID:      0,
		ClientID:    0,
		Expires:      time.Now().UTC().Add(time.Hour * expiration).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}