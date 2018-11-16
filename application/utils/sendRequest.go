package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func SendRequest(url string, method string, jsonStr []byte) ([]byte, int, error) {

	client := &http.Client{}
	var req *http.Request
	var err error

	req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth("HttpService", "HttpService")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	var bodyText []byte

	if err != nil {
		bodyText = []byte(err.Error())
		return bodyText, 400, err
	}

	bodyText, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		bodyText = []byte(err.Error())
		return bodyText, 400, err
	}

	return bodyText, resp.StatusCode, nil
}