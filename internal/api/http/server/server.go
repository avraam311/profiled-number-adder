package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/avraam311/profiled-number-adder/internal/api/http/handlers/numbers"
)

func NewRouter(handlerNums *numbers.Handler) *gin.Engine {
	e := gin.Default()

	numsGroup := e.Group("/numbers")
	{
		numsGroup.POST("/add-up", handlerNums.AddUp)
	}

	debugGroup := e.Group("/debug/pprof")
	pprof.RouteRegister(debugGroup, "")

	return e
}

func NewServer(addr string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}
