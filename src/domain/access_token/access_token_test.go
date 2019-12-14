package access_token

import "testing"

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired(){
		t.Error("brand new access token should not be expired")
	}
	if at.AccessToken != ""{
		t.Error("brand new access token should not have defined access token id")
	}
	if at.UserID != 0{
		t.Error("brand new access token should not have define user id")
	}
}
