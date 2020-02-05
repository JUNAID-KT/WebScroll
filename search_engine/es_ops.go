package search_engine

import (
	"encoding/json"

	"github.com/JUNAID-KT/WebScroll/models"
	"github.com/JUNAID-KT/WebScroll/util"
	log "github.com/Sirupsen/logrus"
	"github.com/olivere/elastic"
)

func (es *esEngine) SaveWebContent(doc models.Website) error {
	_, err := es.Client.
		Index().
		Index(util.WebScrapIndexName).
		Type(util.WebScrapTypeName).
		BodyJson(doc).
		Do(es.Ctx)

	if err != nil {
		log.WithFields(log.Fields{"method": "SaveWebContent", "Index Name": util.WebScrapIndexName,
			"error": err.Error()}).
			Error("error occurred while saving web content")
		return err
	}

	return nil
}

// Search in DB ; matching for the given text input against the content
func (es *esEngine) GetURL(text string) (error, string) {
	match_query := elastic.NewQueryStringQuery(text)
	searchResult, err := es.Client.Search().
		Index(util.WebScrapIndexName).
		Query(match_query.Field("content")).
		Size(1).
		Do(es.Ctx)
	if err != nil {
		return err, ""
	}
	var url string
	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index
			var websites models.Website
			// Deserialize hit.Source
			err := json.Unmarshal(*hit.Source, &websites)
			if err != nil {
				// Deserialization failed
				return err, ""
			}
			url = websites.URL
		}
	} else {
		// No hits
		return nil, ""
	}
	return nil, url
}
