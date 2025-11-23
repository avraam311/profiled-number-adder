package server

import (
	"net/http"

	"github.com/avraam311/profiled-number-adder/internal/api/http/handlers/numbers"
	"github.com/gin-gonic/gin"
)

func NewRouter(handlerNums *numbers.Handler) *gin.Engine {
	e := gin.Default()

	numsGroup := e.Group("/numbers")
	{
		numsGroup.POST("/add-up", handlerNums.AddUp)
	}

	return e
}

func NewServer(addr string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
