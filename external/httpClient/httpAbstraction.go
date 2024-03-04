package httpClient

type HTTPClient interface {
	Get(url string, headers map[string]string) ([]byte, error)
	Post(url string, data []byte, headers map[string]string) ([]byte, error)
	// Add other necessary methods (PUT, DELETE, etc.)
}
