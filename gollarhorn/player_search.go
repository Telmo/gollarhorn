package gollarhorn

import "fmt"

func (ps *PlayerService) GetMembershipId(platType string, playerName string) (string, error) {
	plat := platforms[platType]
	u := fmt.Sprintf("SearchDestinyPlayer/%v/%v/", plat, playerName)

	r, _, err := ps.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return "", err
	}
	//todo: somthing with that shit
	return r.Response.([]interface{})[0].(map[string]interface{})["membershipId"].(string), nil
}
