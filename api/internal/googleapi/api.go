package googleapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Piokor/olutek_lib/internal"
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

func GetBook(bookId string) (*internal.Book, error) {
	url, err := getDetailUrl(bookId)
	if err != nil {
		return nil, err
	}
	responseBody, err := makeRequest(url)
	if err != nil {
		return nil, err
	}

	responseVolume := googleVolume{}
	err = json.Unmarshal(responseBody, &responseVolume)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling Google API http response: %s", err)
	}
	return responseVolume.toBook(), nil
}

func GetBooks(query *BookQuery) ([]*internal.Book, error) {
	url, err := getListUrl(query)
	if err != nil {
		return nil, err
	}
	responseBody, err := makeRequest(url)
	if err != nil {
		return nil, err
	}

	responseBooks := googleVolumeResults{}
	err = json.Unmarshal(responseBody, &responseBooks)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling Google API http response: %s", err)
	}

	result := make([]*internal.Book, len(responseBooks.Items))
	for i, gv := range responseBooks.Items {
		result[i] = gv.toBook()
	}
	return result, nil
}
