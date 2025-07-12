package utils

import (
	"io"
	"net/http"
	"os"
)

func NewRequestSIM(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("SIM_API_KEY_SERVICE"))
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func NewRequestGuest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	return req, nil
}
