package endpoint

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Endpoint struct {
}

var (
	upgrade websocket.Upgrader
)

func init() {
	upgrade = websocket.Upgrader{}
}
func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (e Endpoint) NotFound(context *gin.Context) {
	context.JSON(http.StatusNotFound, struct {
		Time        time.Time `json:"time"`
		Description string    `json:"description"`
	}{
		Time:        time.Now(),
		Description: "Workers are working, coming soon!!!",
	})
}

func (e Endpoint) Websocket(ctx *gin.Context) {
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
