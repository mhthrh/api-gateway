package endpoint

import (
	"api-gateway/pkg/service"
	"api-gateway/pkg/transport/customer/decode"
	"api-gateway/pkg/transport/customer/encode"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	d decode.Decode
	e encode.Encode
	g service.Grpc
}

func (e Endpoint) CustomerRegister(c *gin.Context) {
	customer, err := e.d.RegisterCustomer(c.Request.Body)
	fmt.Println(customer, err)
	a, b := e.g.Register()
	fmt.Println(a, b)
	code, res := e.e.RegisterCustomer(a, nil)
	c.JSON(code, res)
	c.Done()
	// write response
}
