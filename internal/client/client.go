package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	httpProtocol = "http://"
	wsProtocol   = "ws://"
	domain       = "localhost:80"
	apiEndpoint  = httpProtocol + domain
	wsEndpoint   = wsProtocol + domain
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func sendRequest(method string, url string, rb io.Reader) ([]byte, error) {
	httpClient := http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, rb)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request failed with %d code", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
