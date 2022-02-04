package app

import (
	"github.com/gin-gonic/gin"
	"github.com/toan267/bookstore_oauth-api/src/clients/cassandra"
	"github.com/toan267/bookstore_oauth-api/src/domain/access_token"
	"github.com/toan267/bookstore_oauth-api/src/http"
	"github.com/toan267/bookstore_oauth-api/src/repository/db"
)
var (
	router = gin.Default()
)
func StartApplication() {
	//session, dbErr := cassandra.GetSession()
	//if dbErr != nil {
	//	panic(dbErr)
	//}
	//defer session.Close()
	defer cassandra.GetSession().Close()
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/accecss_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/accecss_token", atHandler.Create)
	router.Run(":8080")
}