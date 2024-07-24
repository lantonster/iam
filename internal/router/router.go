package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/ginkit"
	"github.com/lantonster/iam/config"
	_ "github.com/lantonster/iam/docs/api"
	"github.com/lantonster/iam/internal/handler"
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

func NewRouter(conf *config.Config, h *handler.Handler) http.Handler {
	engine := gin.Default()

	// middleware
	engine.Use(ginkit.CorsMiddleware)

	// swagger
	defer fmt.Printf("API docs: http://localhost:%d/swagger/index.html\n", conf.Port)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router
	initAuthRouter(engine, h.Auth)

	return engine
}
