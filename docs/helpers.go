package data_parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type APIError struct {
	Code    int
	Message string
}

func APIErrorFromResponse(response *http.Response) *APIError {
	var apiError APIError

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&apiError)
	if err!= nil {
		log.Println("Error parsing API response:", err)
		return nil
	}

	return &apiError
}

func GetAPIResponse(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err!= nil {
		return nil, err
	}

	if resp.StatusCode!= http.StatusOK {
		return nil, fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	return resp, nil
}

func ParseJSONResponse(response *http.Response) (interface{}, error) {
	decoder := json.NewDecoder(response.Body)
	var data interface{}
	err := decoder.Decode(&data)
	if err!= nil {
		return nil, err
	}

	return data, nil
}

func TrimWhitespace(s string) string {
	return strings.Trim(s, " \t\n\r")
}

func IsEmpty(s string) bool {
	return s == ""
}

func IsEmptyStringSlice(s []string) bool {
	return len(s) == 0
}

func IsNil(v interface{}) bool {
	return v == nil
}

func IsZero(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return v == 0
	case string:
		return v == ""
	case bool:
		return!v.(bool)
	default:
		return false
	}
}