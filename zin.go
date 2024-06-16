package zin

import (
	"github.com/devisettymahidhar315/zin/api"
	"github.com/gin-gonic/gin"
)

func Zin() *gin.Engine {

	// Initialize the Gin router and routes
	return api.InitializeRoutes()

}
