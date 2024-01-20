package service

import v1 "credit/pkg/service/v1"

type ServiceLayer interface {
	v1.ServiceLayer
}

func NewService() ServiceLayer {
	return &srvSt{v1.NewV1Service()}
}

type srvSt struct {
	v1.ServiceLayer
}
