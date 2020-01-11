package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/IkezawaYuki/videostore_oauth-api/src/domain/users"
	"github.com/IkezawaYuki/videostore_utils-go/rest_errors"
)

var (
	// usersRestClient = rest.RequestBuilder{
	// 	BaseURL: "https://api.videostore.com",
	// 	Timeout: 100 * time.Millisecond,
	// }
	httpClient = &http.Client{
		Timeout: 100 * time.Millisecond,
	}
)

// UsersRepository is ...
type UsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct {
}

// NewRestUsersRepository is ...
func NewRestUsersRepository() UsersRepository {
	return &usersRepository{}
}

// NewRepository is ...
func NewRepository() UsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	bytes, _ := json.Marshal(request)
	fmt.Println(string(bytes))

	// response := usersRestClient.Post("/users/login", request)
	u := url.Values{}
	u.Add("email", email)
	u.Add("password", password)
	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.videostore.com/users/login",
		strings.NewReader(u.Encode()),
	)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("invalid response when trying to login user", errors.New("restclient error"))
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("invalid response when trying to login user", errors.New("restclient error"))
	}
	// if response == nil || response.Response == nil {
	// 	return nil, rest_errors.NewInternalServerError("invalid response when trying to login user", errors.New("restclient error"))
	// }
	if response.StatusCode > 299 {
		body, err := ioutil.ReadAll(response.Body)
		apiErr, err := rest_errors.NewRestErrorFromBytes(body)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr
	}
	var user users.User
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users response", errors.New("json parsing error"))
	}
	return &user, nil
}
