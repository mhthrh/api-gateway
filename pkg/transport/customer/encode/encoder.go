package encode

import (
	_ "github.com/mhthrh/GoNest/model/customer"
	cError "github.com/mhthrh/GoNest/model/error"
	http "golang.org/x/net/http2"
)

type Encode struct {
}

func (d Encode) RegisterCustomer(res interface{}, e *cError.XError) (code http.ErrCode, json string) {
	return 0, ""
}
