package twitchappapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const apiURL = "https://addons-ecs.forgesvc.net/api/v2"

var client = &http.Client{
	Timeout: time.Second * 10,
}

func requestData(request *http.Request) ([]byte, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request responded with status code %d, not 200 OK", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

func requestPlainText(method, url string, body []byte) (string, error) {
	var plainText string

	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return plainText, err
	}

	request.Header.Add("Content-Type", "text/plain")

	data, err := requestData(request)
	if err != nil {
		return plainText, err
	}

	return string(data), nil
}

func requestJSONData(method, url string, body []byte, v interface{}) error {
	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")

	data, err := requestData(request)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}
