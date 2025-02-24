package main

import (
	"api-gateway/pkg/transport"
	"context"
	"fmt"
	cError "github.com/mhthrh/GoNest/model/error"
	loader "github.com/mhthrh/GoNest/pkg/loader/file"
	l "github.com/mhthrh/GoNest/pkg/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	configPath   = "/customer-service/config"
	configName   = "config.json"
	readTimeOut  = 10 * time.Second
	WriteTimeOut = 10 * time.Second
	idleTimeOut  = 180 * time.Second
)

var (
	osInterrupt       chan os.Signal
	internalInterrupt chan *cError.XError
)

func init() {
	osInterrupt = make(chan os.Signal)
	internalInterrupt = make(chan *cError.XError)
	signal.Notify(osInterrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP)
}
func main() {

	ctx, cancel := context.WithCancel(context.Background())
	logger := zap.New(l.LogConfig())
	defer func() {
		_ = logger.Sync()
	}()

	sugar := logger.Sugar()
	sugar.Info("Loading config...")

	config, err := loader.New(configPath, configName).Initialize()
	if err != nil {
		sugar.Fatal(err)
	}
	sugar.Info("customer.v1 service config loaded successfully")
	t := transport.Transport{}
	httpRest := http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.HTTP.Ip, config.HTTP.Port),
		Handler:      t.Http(ctx),
		TLSConfig:    nil,
		ReadTimeout:  readTimeOut,
		WriteTimeout: WriteTimeOut,
		IdleTimeout:  idleTimeOut,
	}
	socket := http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.WebSocket.Ip, config.WebSocket.Port),
		Handler:      t.WebSkt(ctx),
		TLSConfig:    nil,
		ReadTimeout:  readTimeOut,
		WriteTimeout: WriteTimeOut,
		IdleTimeout:  idleTimeOut,
	}
	go func() {
		if e := httpRest.ListenAndServe(); e != nil {
			log.Fatalf("failed to serve: %v \n", err)
			return
		}
	}()
	go func() {
		if e := socket.ListenAndServe(); e != nil {
			log.Fatalf("failed to serve: %v \n", err)
			return
		}
	}()
	select {
	case <-osInterrupt:
		sugar.Info("OS interrupt signal received")
	case e := <-internalInterrupt:
		sugar.Errorf("customer.v1 service listener interrupt signal received, %+v", e)
	}

	sugar.Info("stopping customer.v1 service...")
	cancel()

	<-internalInterrupt
}
