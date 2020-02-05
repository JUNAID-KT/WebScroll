package models

type Status struct {
	Status StatusResponse `json:"status"`
}

type StatusResponse struct {
	StatusCode      int    `json:"status_code"`
	DescriptionCode string `json:"description_code"`
	Description     string `json:"description"`
}
type UrlScrapRequest struct {
	URL string `json:"url"`
}
type SearchRequest struct {
	Text string `json:"text_input"`
}
type SearchResponse struct {
	Status StatusResponse `json:"status"`
	URL    string         `json:"url"`
}
