package decode

import (
	"github.com/mhthrh/GoNest/model/customer"
	cError "github.com/mhthrh/GoNest/model/error"
)

type Decode struct {
}

func (d Decode) RegisterCustomer(req interface{}) (*customer.Customer, *cError.XError) {
	return nil, nil
}
