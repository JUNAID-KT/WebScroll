package controller

import (
	"net/http"

	"github.com/JUNAID-KT/WebScroll/models"
	se "github.com/JUNAID-KT/WebScroll/search_engine"
	"github.com/JUNAID-KT/WebScroll/util"
	"github.com/gin-gonic/gin"
)

var GetUrl = "GetURL"

// Handler for getting transactions
func GetURL(context *gin.Context) {
	var apiRequest models.SearchRequest
	if err := context.BindJSON(&apiRequest); err != nil {
		util.ErrorResponder(err, GetUrl, util.FailureDesc, util.BindingFailedMsg, http.StatusBadRequest, context)
		return
	}
	es := se.GetESInstance()
	err, url := es.GetURL(apiRequest.Text)
	if err != nil {
		util.ErrorResponder(err, GetUrl, util.FailureDesc, err.Error(), http.StatusInternalServerError, context)
		return
	}
	context.JSON(http.StatusOK, models.SearchResponse{
		Status: util.SetStatusResponse(http.StatusOK,
			http.StatusText(http.StatusOK),
			"URL fetched"),
		URL: url,
	})
}
