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
		WebScroll := v1.Group(util.ApiPrefix)
		{
			// routes
			WebScroll.POST(util.SearchURL, controller.GetURL)
			WebScroll.POST(util.ScrapURL, controller.ScrapWeb)
			WebScroll.GET(util.Index, controller.RenderIndexPage)

		}
	}
	return router
}
