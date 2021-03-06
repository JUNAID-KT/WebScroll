package router

import (
	"net/http"

	"github.com/JUNAID-KT/WebScroll/util"

	"github.com/gin-gonic/gin"
)

// InitRoutes : registers all routers for the application.
func InitRoutes() *gin.Engine {
	// Start HTTP server
	router := gin.Default()
	router.LoadHTMLGlob("./template/*")
	router = SetRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,
			util.SetStatus(http.StatusNotFound,
				util.FailureDesc,
				"Invalid Request"))
	})
	return router
}
