package api

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

func (S *ApiData) Initiate(config map[string]string, logger *logrus.Logger) {
	S.Client = &http.Client{}
	S.URL = config["API"]

}

func (A *ApiData) GetSingleData(logger *logrus.Logger) (string, bool) {
	var (
		client = A.Client
	)

	resp, err := client.Get(A.URL)
	if err != nil {
		logger.Error(err, " status code: ", resp.StatusCode)
		return "", false
	}
	defer resp.Body.Close()
	word, err := io.ReadAll(resp.Body)

	if err != nil {
		logger.Error(err)
		return "", false
	}
	if len(word) <= 0 {
		logger.Error("Api response is empty")
		return "", false
	}

	resultString := CutOff(string(word))

	return resultString, true

}

func CutOff(oldString string) string {

	replacer := strings.NewReplacer("[", "", "]", "", "\"", "")

	return replacer.Replace(oldString)

}
