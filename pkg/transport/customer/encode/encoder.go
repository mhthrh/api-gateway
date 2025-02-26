package encode

import (
	_ "github.com/mhthrh/GoNest/model/customer"
	cError "github.com/mhthrh/GoNest/model/error"
)

type Encode struct {
}

func (d Encode) RegisterCustomer(res interface{}, e *cError.XError) (code int, json string) {
	return 0, ""
}
