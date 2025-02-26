package transport

import (
	"api-gateway/pkg/endpoint"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	trs Trans
	e   endpoint.Endpoint
)

func init() {
	gin.SetMode(gin.TestMode)
	gin.DisableConsoleColor()
	trs = Trans{}
	e = endpoint.Endpoint{}
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
	cGroup := r.Group("customer")

	cGroup.POST("/create", e.CustomerRegister)

	r.NoRoute(trs.NotFound)
	return r
}

func (t Transport) WebSkt(ctx context.Context) http.Handler {
	router := gin.New()

	router.Use(gin.Recovery())
	router.GET("/x-bank/ws", trs.Websocket)
	router.NoRoute(trs.NotFound)

	return router
}
