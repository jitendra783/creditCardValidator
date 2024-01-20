package api

import (
	v1 "credit/pkg/service/v1"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	fmt.Println("inside router")
	router := gin.Default()
	router.GET("/valid", v1.NewV1Service().Validate)
	return router
}
