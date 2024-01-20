package src

import (
	"credit/pkg/service/v1/creditCard/repo"

	"github.com/gin-gonic/gin"
)

type SrcLayer interface {
	Validate(c *gin.Context)
}

func NewSrcService() SrcLayer {
	return &srvSt{r: repo.NewRepoService()}
}

type srvSt struct {
    r repo.RepoLayer
}
