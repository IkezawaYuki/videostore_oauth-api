package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {

}
func TestLoginUserInvalidErrorInterface(t *testing.T) {

}
func TestLoginUserInvalidLoginCredentials(t *testing.T) {

}
func TestLoginUserInvalidUserJsonResponse(t *testing.T) {

}
func TestLoginUserNoError(t *testing.T) {

}
