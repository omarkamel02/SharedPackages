package httpClient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var (
	_ErrServerError = "error while connecting to server"
	//_               abstractions.HTTPClient = &StandardHTTPClient{} // force implementation of the interface
)

type StandardHTTPClient struct {
}

func NewStandardHTTPClient() StandardHTTPClient {
	return StandardHTTPClient{}
}

// Get implements abstractions.HTTPClient.
func (StandardHTTPClient) Get(url string, headers map[string]string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(_ErrServerError, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Post implements abstractions.HTTPClient.
func (StandardHTTPClient) Post(url string, data []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(_ErrServerError, resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
