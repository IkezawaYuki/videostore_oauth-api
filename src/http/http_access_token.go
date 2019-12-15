package http

import (
	"github.com/IkezawaYuki/videostore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler{
	return &accessTokenHandler{
		service:service,
	}
}

func (handler *accessTokenHandler) GetByID(c *gin.Context){
	accessTokenID := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.service.GetByID(accessTokenID)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}