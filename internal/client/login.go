package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	loginPath = "/user/login"
)

type LoginResponce struct {
	Url string
}

func Login(n int) (result []string, err error) {

	for i := 1; i <= n; i++ {
		body, _ := json.Marshal(map[string]string{
			"userName": fmt.Sprintf("testUser%v", i),
			"password": fmt.Sprintf("password%v", i),
		})

		rb := bytes.NewBuffer(body)

		resp, err := sendLoginRequest(rb)
		if err != nil {
			return nil, err
		}

		result = append(result, extractToken(resp.Url))
	}
	return
}

func sendLoginRequest(rb io.Reader) (*LoginResponce, error) {
	url := apiEndpoint + loginPath

	body, err := sendRequest(http.MethodPost, url, rb)
	if err != nil {
		return nil, err
	}

	result := LoginResponce{}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("сan not unmarshal JSON")
		return nil, err
	}

	return &result, nil
}

func extractToken(url string) (token string) {
	if arr := strings.SplitAfter(url, "="); len(arr) > 0 {
		token = arr[1]
	}
	return
}
