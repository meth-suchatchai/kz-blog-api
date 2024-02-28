package kzline

import (
	"encoding/json"
	"fmt"
)

type StatusResponse struct {
	Status     int    `json:"status"`
	Message    string `json:"message"`
	TargetType string `json:"targetType"`
	Target     string `json:"target"`
}

func (c *defaultClient) GetApiStatus() (*StatusResponse, error) {
	url := fmt.Sprintf("%s/api/status", c.LineApi)
	response, err := c.restyClient.R().
		SetHeader("Authorization", c.getHeaderToken()).
		Get(url)
	if err != nil {
		return nil, err
	}

	result := StatusResponse{}

	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
