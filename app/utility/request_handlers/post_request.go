package request_handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PostRequest(uri string, data interface{}) (string, error) {

	// json encode interface
	b, err := json.Marshal(data)
	var jsonStr = []byte(b)

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	body, bodyError := ioutil.ReadAll(resp.Body)
	if bodyError != nil {
		return "", bodyError
	}

	return string(body), err
}