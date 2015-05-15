package gollarhorn

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version            = "0.1"
	defaulBungieURL    = "http://www.bungie.net/"
	defaultPlatformURL = defaulBungieURL + "Platform/Destiny/"
	defaultUserAgent   = "gollarhorn/" + version
)

type ResponseData map[string]interface{}

var platforms = map[string]int{
	"xbox":   1,
	"psn":    2,
	"bungie": 254,
}

type Client struct {
	client      *http.Client
	BungieURL   *url.URL
	PlatformURL *url.URL
	UserAgent   string
	Character   *CharacterService
	Player      *PlayerService
	Platform    *PlatformService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	bungieURL, _ := url.Parse(defaulBungieURL)
	platURL, _ := url.Parse(defaultPlatformURL)

	c := &Client{
		client:      httpClient,
		BungieURL:   bungieURL,
		PlatformURL: platURL,
		UserAgent:   defaultUserAgent,
	}

	c.Platform = &PlatformService{client: c}
	c.Character = &CharacterService{client: c}
	c.Player = &PlayerService{client: c}
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.PlatformURL.ResolveReference(rel)

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

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (cleint *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := cleint.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
