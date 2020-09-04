package http_client

import "github.com/go-resty/resty/v2"

// HTTPClient represents a client uses Discourse API.
type HTTPClient struct {
	client *resty.Client
}

// NewHTTPClient contruct
func NewHTTPClient(host string, apiUser string, apiKey string) *HTTPClient {
	client := resty.New()
	client.SetHeader("Accept", "application/json")
	client.SetHostURL(host)

	return &HTTPClient{client: client}
}

// Client returns the name of host.
func (h *HTTPClient) Client() *resty.Client {
	return h.client
}

// Get send a request with GET method to Discourse and returns the result.
func (h *HTTPClient) Get(endpoint string, params map[string]string) []byte {
	response, err := h.client.R().
		SetQueryParams(params).
		Get(endpoint)

	if err != nil {
		return []byte{}
	}

	return response.Body()
}
