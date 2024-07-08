package googleapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Piokor/olutek_lib/internal/bookapi"
)

func makeRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making Google API http request: %s", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetVolume(volumeId string) (*bookapi.Volume, error) {
	url, err := GetDetailUrl(volumeId)
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
	volume := responseVolume.ToVolume()
	return &volume, nil
}

func GetVolumes(query string) ([]bookapi.Volume, error) {
	url, err := GetListUrl(query)
	if err != nil {
		return nil, err
	}
	responseBody, err := makeRequest(url)
	if err != nil {
		return nil, err
	}

	responseVolumes := GoogleVolumeResults{}
	err = json.Unmarshal(responseBody, &responseVolumes)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling Google API http response: %s", err)
	}

	result := make([]bookapi.Volume, len(responseVolumes.Items))
	for i, gv := range responseVolumes.Items {
		result[i] = gv.ToVolume()
	}
	return result, nil
}
