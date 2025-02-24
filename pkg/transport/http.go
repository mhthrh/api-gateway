package transport

import (
	"api-gateway/pkg/endpoint"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	e *endpoint.Endpoint
)

func init() {
	gin.SetMode(gin.TestMode)
	gin.DisableConsoleColor()
	e = endpoint.NewEndpoint()
}

type Transport struct {
}

func New() ITransport {
	return Transport{}
}

func (t Transport) Http(ctx context.Context) http.Handler {
	r := gin.New()
	//router.Use(Controller.Middleware)
	r.Use(gin.Recovery())

	r.POST("/login", nil)
	r.NoRoute(e.NotFound)
	return r
}

func (t Transport) WebSkt(ctx context.Context) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/x-bank/ws", e.Websocket)
	router.NoRoute(e.NotFound)

	return router
}
