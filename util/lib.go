package util

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/JUNAID-KT/WebScroll/models"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func MakeRequest(URL string) (string, error) {
	// Make request
	response, err := http.Get("https://" + URL)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	// Get the response body as a string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)
	return pageContent, nil
}

//  SetStatus : setting models.Status
func SetStatus(statusCode int, descriptionCode string, description string) models.Status {
	return models.Status{
		Status: models.StatusResponse{
			StatusCode:      statusCode,
			DescriptionCode: descriptionCode,
			Description:     description,
		},
	}
}

//SetStatusResponse: Creates status response
func SetStatusResponse(statusCode int, descriptionCode string, description string) models.StatusResponse {
	return models.StatusResponse{
		StatusCode:      statusCode,
		DescriptionCode: descriptionCode,
		Description:     description,
	}
}

//ErrorResponder: Creates a JSON response of the corresponding error
func ErrorResponder(err error, method string, descrCode string, description string, httpCode int, ctx *gin.Context) {
	if err != nil {
		log.Error(err)
		log.WithFields(log.Fields{"method": method, "description": description, "error": err.Error()}).
			Infoln("error occurred ")
	}
	log.WithFields(log.Fields{"method": method, "description": description}).Errorln("error occurred ")

	ctx.JSON(httpCode, SetStatus(httpCode, descrCode, description))
}
func Retry(attempts int, sleep time.Duration, f func() error) error {
	if err := f(); err != nil {
		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating race condition
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2
			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, f)
		}
		return err
	}
	return nil
}
