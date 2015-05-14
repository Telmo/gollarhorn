package gollarhorn

import (
	"fmt"
	"net/http"
)

type SearchResponse struct {
	PlatformResponse
	Response []map[string]interface{} `json: Response`
}

func (ps *PlayerService) SearchPlayer(platform string, playerName string) (*SearchResponse, *http.Response, error) {
	plat := platforms[platform]
	u := fmt.Sprintf("SearchDestinyPlayer/%v/%v/", plat, playerName)

	req, err := ps.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(SearchResponse)
	resp, err := ps.client.Do(req, r)

	if err != nil {
		return nil, nil, err
	}

	return r, resp, nil
}

func (ps *PlayerService) GetMembershipId(platform string, playerName string) (string, error) {
	r, _, err := ps.SearchPlayer(platform, playerName)
	if err != nil {
		return "", err
	}

	return r.Response[0]["membershipId"].(string), nil
}
