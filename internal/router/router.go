package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lzaun/iam/config"
	_ "github.com/lzaun/iam/docs/api"
	"github.com/lzaun/iam/internal/router/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title						Iam API
//	@description				authorization
//	@version					1.0
//	@host						localhost:8080
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@securityDefinitions.apikey	apiKey
//	@in							header
//	@name						Authorization

func NewRouter(conf *config.Config) http.Handler {
	engine := gin.Default()

	engine.Use(middleware.Cors)

	defer fmt.Printf("API docs: http://localhost:%d/swagger/index.html\n", conf.Port)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return engine
}
