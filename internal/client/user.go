package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	userPath = "/user"
)

type UserResponce struct {
	ID       string
	UserName string
}

func Create(n int) (result []UserResponce, err error) {
	for i := 1; i <= n; i++ {
		body, _ := json.Marshal(map[string]string{
			"userName": fmt.Sprintf("testUser%v", i),
			"password": fmt.Sprintf("password%v", i),
		})

		rb := bytes.NewBuffer(body)

		resp, err := sendCreateRequest(rb)
		if err != nil {
			return nil, err
		}

		result = append(result, UserResponce{
			ID:       resp.ID,
			UserName: resp.UserName,
		})
	}

	return
}

func sendCreateRequest(requestBody io.Reader) (*UserResponce, error) {
	url := apiEndpoint + userPath

	body, err := sendRequest(http.MethodPost, url, requestBody)
	if err != nil {
		return nil, err
	}

	result := UserResponce{}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("сan not unmarshal JSON")
		return nil, err
	}

	return &result, nil
}
