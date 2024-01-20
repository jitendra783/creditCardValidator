package v1

import valid "credit/pkg/service/v1/creditCard"

type ServiceLayer interface {
	valid.ValidLayer
}

func NewV1Service() ServiceLayer {
	return &srvSt{valid.NewValidService()}
}

type srvSt struct {
	valid.ValidLayer
}
