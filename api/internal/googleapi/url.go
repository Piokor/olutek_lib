package googleapi

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const base_url = "https://www.googleapis.com/books/v1/volumes"

func getApiKey() (string, error) {
	apiKey, present := os.LookupEnv("GOOGLE_API_KEY")
	if !present {
		return "", errors.New("google API key not set up")
	}
	return apiKey, nil
}

func getDetailUrl(volumeId string) (string, error) {
	apiKey, err := getApiKey()
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/%s?key=%s", base_url, volumeId, apiKey)
	return url, nil
}

func createOptionalParamsString(query *BookQuery) string {
	paramsMap := map[string]string{
		"intitle":     query.Title,
		"inauthor":    query.Author,
		"inpublisher": query.Publisher,
		"isbn":        query.Isbn,
	}
	optionalParams := make([]string, 0, 4)
	for param, value := range paramsMap {
		if value == "" {
			continue
		}
		urlParam := fmt.Sprintf("+%s:%s", param, value)
		optionalParams = append(optionalParams, urlParam)
	}
	return strings.Join(optionalParams, "")
}

func createQueryString(query *BookQuery) string {
	optionalParamsStr := createOptionalParamsString(query)
	return fmt.Sprintf("q=%s%s", query.Query, optionalParamsStr)
}

func getListUrl(query *BookQuery) (string, error) {
	apiKey, err := getApiKey()
	keyQuery := fmt.Sprintf("key=%s", apiKey)
	if err != nil {
		return "", err
	}
	queryStr := createQueryString(query)
	escapedQuery := url.PathEscape(queryStr)
	url := fmt.Sprintf("%s?%s&%s", base_url, keyQuery, escapedQuery)
	return url, nil
}
