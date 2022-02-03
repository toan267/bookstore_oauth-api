package app

import (
	"github.com/gin-gonic/gin"
	"github.com/toan267/bookstore_oauth-api/src/domain/access_token"
	"github.com/toan267/bookstore_oauth-api/src/http"
	"github.com/toan267/bookstore_oauth-api/src/repository/db"
)
var (
	router = gin.Default()
)
func StartApplication() {
	//dbRepository := db.New()
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/accecss_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}