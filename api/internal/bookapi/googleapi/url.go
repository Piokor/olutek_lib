package googleapi

import (
	"errors"
	"fmt"
	"net/url"
	"os"
)

const BASE_URL = "https://www.googleapis.com/books/v1/volumes"

func getApiKey() (string, error) {
	apiKey, present := os.LookupEnv("GOOGLE_API_KEY")
	if !present {
		return "", errors.New("google API key not set up")
	}
	return apiKey, nil
}

func GetDetailUrl(volumeId string) (string, error) {
	apiKey, err := getApiKey()
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/%s?key=%s", BASE_URL, volumeId, apiKey)
	return url, nil
}

func GetListUrl(query string) (string, error) {
	apiKey, err := getApiKey()
	if err != nil {
		return "", err
	}
	escapedQuery := url.PathEscape(query)
	url := fmt.Sprintf("%s?key=%s&q=%s", BASE_URL, apiKey, escapedQuery)
	return url, nil
}
