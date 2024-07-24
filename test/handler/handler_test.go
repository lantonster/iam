package handler

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lantonster/iam/config"
	"github.com/lantonster/iam/internal/handler"
	"github.com/lantonster/iam/internal/router"
	"github.com/lantonster/iam/test/service"
)

var (
	h http.Handler
	c *gin.Context

	mockService     *service.MockService
	mockAuthService *service.MockAuthService
)

func setupMock(ctrl *gomock.Controller) {
	mockService = service.NewMockService(ctrl)
	mockAuthService = service.NewMockAuthService(ctrl)

	c = &gin.Context{}
	h = router.NewRouter(config.NewConfig(), handler.NewHandler(mockService))
}

func doRequest(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	return w
}
