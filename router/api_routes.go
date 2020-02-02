package router

import (
	"github.com/JUNAID-KT/WebScroll/controller"

	"github.com/JUNAID-KT/WebScroll/util"
	"github.com/gin-gonic/gin"
)

// Set routers
func SetRoutes(router *gin.Engine) *gin.Engine {

	v1 := router.Group(util.ApiVersion)
	{
		wallet := v1.Group(util.ApiPrefix)
		{
			// routes
			wallet.GET(util.SearchURL, controller.GetURL)
		}
	}
	return router
}
