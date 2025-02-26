package transport

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	cTransport "github.com/mhthrh/GoNest/model/transport"
	convert "github.com/mhthrh/GoNest/pkg/util/convertor"
	"log"
	"net/http"
)

type Trans struct {
}

var (
	upgrade websocket.Upgrader
)

func init() {
	upgrade = websocket.Upgrader{}
}
func NewEndpoint() *Trans {
	return &Trans{}
}
func (e Trans) AccessDenied(context *gin.Context) {
	j := convert.Json{}
	res, _ := j.Marshal(*cTransport.ForbiddenError(nil))
	context.AbortWithStatusJSON(http.StatusForbidden, res)
}
func (e Trans) NotFound(context *gin.Context) {
	j := convert.Json{}
	res, _ := j.Marshal(*cTransport.NotFoundError(nil))
	context.AbortWithStatusJSON(http.StatusNotFound, res)
}

func (e Trans) Websocket(ctx *gin.Context) {
	toWs := make(chan string)
	fromWs := make(chan string)
	cntx, cancel := context.WithCancel(context.Background())
	defer cancel()
	w, r := ctx.Writer, ctx.Request

	//check authentication
	upgrade.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	wSc, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		ctx.JSON(http.StatusUpgradeRequired, "connote upgrade connection to websocket")
		return
	}
	go read(cntx, wSc, fromWs)
	go write(cntx, wSc, toWs)

}

func read(ctx context.Context, c *websocket.Conn, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("read: %v", err)
				return
			}
			ch <- string(message)
		}
	}
}
func write(ctx context.Context, c *websocket.Conn, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-ch:
			//string response type is 1
			err := c.WriteMessage(1, []byte(msg))
			if err != nil {
				log.Printf("read: %v", err)
				return
			}
		}
	}

}
