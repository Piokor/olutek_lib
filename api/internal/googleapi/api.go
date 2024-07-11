package googleapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func makeRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making Google API http request: %s", err)
	}
	defer res.Body.Close()
	status := res.StatusCode
	if status != 200 {
		return nil, fmt.Errorf("error making Google API http request, status code: %d", status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetBook(bookId string) (*GoogleVolume, error) {
	url, err := getDetailUrl(bookId)
	if err != nil {
		return nil, err
	}
	responseBody, err := makeRequest(url)
	if err != nil {
		return nil, err
	}

	responseVolume := GoogleVolume{}
	err = json.Unmarshal(responseBody, &responseVolume)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling Google API http response: %s", err)
	}
	return &responseVolume, nil
}

func GetBooks(query *BookQuery) ([]*GoogleVolume, error) {
	url, err := getListUrl(query)
	if err != nil {
		return nil, err
	}
	responseBody, err := makeRequest(url)
	if err != nil {
		return nil, err
	}
	responseBooks := GoogleVolumeResults{}
	err = json.Unmarshal(responseBody, &responseBooks)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling Google API http response: %s", err)
	}
	result := make([]*GoogleVolume, len(responseBooks.Items))
	for i, gv := range responseBooks.Items {
		result[i] = &gv
	}
	return result, nil
}
