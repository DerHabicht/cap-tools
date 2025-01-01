package router

import (
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) *gin.RouterGroup {
	v1 := engine.Group("/v1")

	return v1
}
