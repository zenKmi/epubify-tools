package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zenKmi/epubify-tools/cmd/api/endpoints"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// general endpoints

	// other endpoints
	endpoints.SetupFileAPIRoutes(router)
	return router
}
