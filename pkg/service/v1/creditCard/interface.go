package valid

import (
	"credit/pkg/service/v1/creditCard/repo"
	"credit/pkg/service/v1/creditCard/src"
)

type ValidLayer interface {
	src.SrcLayer
	repo.RepoLayer
}

func NewValidService() ValidLayer {
	return &srvSt{SrcLayer: src.NewSrcService(),
		RepoLayer: repo.NewRepoService()}
}

type srvSt struct {
	src.SrcLayer
	repo.RepoLayer
}
