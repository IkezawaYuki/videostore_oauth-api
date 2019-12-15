package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "brand new access token should not have defined access token id")
	assert.True(t, at.UserID == 0, "brand new access token should not have define user id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "brand new access token should not be expired")

	at.Expires = time.Now().UTC().Add(time.Hour * 3).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should not be expired")
}