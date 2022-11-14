package api

import "net/http"

type ApiData struct {
	Client *http.Client
	URL    string
	Data   string
}

func NewApiData() *ApiData {
	return &ApiData{}
}
