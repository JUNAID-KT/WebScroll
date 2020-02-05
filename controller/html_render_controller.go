package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderIndexPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}
