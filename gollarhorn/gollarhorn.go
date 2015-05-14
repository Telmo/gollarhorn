package gollarhorn

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	version          = "0.1"
	defaultBaseURL   = "http://www.bungie.net/Platform/Destiny/"
	defaultUserAgent = "gollarhorn/" + version
)

type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string
	Character CharacterService
}

func NewClient(httpClient *http.Client) Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL := url.Parse(defaultBaseURL)

	c := Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: defaultUserAgent,
	}
	c.Character = &CharacterService{client: c}
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaTypeV3)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (cleint *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := cleint.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
