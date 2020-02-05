package controller

import (
	"net/http"

	"github.com/JUNAID-KT/WebScroll/models"
	se "github.com/JUNAID-KT/WebScroll/search_engine"
	"github.com/JUNAID-KT/WebScroll/util"
	"github.com/gin-gonic/gin"
)

const WebScrapHandler = "ScrapWeb"

// Handler for scraping url
func ScrapWeb(context *gin.Context) {
	var apiRequest models.UrlScrapRequest
	if err := context.BindJSON(&apiRequest); err != nil {
		util.ErrorResponder(err, WebScrapHandler, util.FailureDesc, util.BindingFailedMsg, http.StatusBadRequest, context)
		return
	}
	webContent, err := util.MakeRequest(apiRequest.URL)
	if err != nil {
		util.ErrorResponder(err, WebScrapHandler, util.FailureDesc, err.Error(), http.StatusInternalServerError, context)
		return
	}
	es := se.GetESInstance()
	err = es.SaveWebContent(models.Website{URL: apiRequest.URL, Content: webContent})
	if err != nil {
		util.ErrorResponder(err, WebScrapHandler, util.FailureDesc, err.Error(), http.StatusInternalServerError, context)
		return
	}
	context.JSON(http.StatusOK, util.SetStatusResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		"Web content saved"))
}
