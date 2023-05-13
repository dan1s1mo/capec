package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func MakeHttpRequest[T any](
	url string,
	bodyString string,
	requestType string,
	headers map[string]string,
) *T {
	data := []byte(bodyString)
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for key, header := range headers {
		req.Header.Set(key, header)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	var responseData T
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &responseData
}
