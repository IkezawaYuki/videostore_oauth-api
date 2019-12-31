package app

import (
	"github.com/IkezawaYuki/videostore_oauth-api/src/clients/cassandra"
	"github.com/IkezawaYuki/videostore_oauth-api/src/http"
	"github.com/IkezawaYuki/videostore_oauth-api/src/repository/db"
	"github.com/IkezawaYuki/videostore_oauth-api/src/repository/rest"
	access_token2 "github.com/IkezawaYuki/videostore_oauth-api/src/service/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session := cassandra.GetSession()
	defer session.Close()
	atHandler := http.NewAccessTokenHandler(access_token2.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token/:access_token_id", atHandler.Create)

	router.Run(":8080")
}
