package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IkezawaYuki/videostore_oauth-api/src/domain/users"
	"github.com/IkezawaYuki/videostore_utils-go/rest_errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.videostore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct {
}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	bytes, _ := json.Marshal(request)
	fmt.Println(string(bytes))

	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid response when trying to login user", errors.New("restclient error"))
	}
	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users response", errors.New("json parsing error"))
	}
	return &user, nil
}
