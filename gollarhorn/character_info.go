package gollarhorn

import "fmt"

func (ps *CharacterService) GetAccountInfo(platType, memshipId string) (map[string]interface{}, error) {
	plat := platforms[platType]
	u := fmt.Sprintf("%v/Account/%v/", plat, memshipId)

	r, _, err := ps.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return nil, err
	}

	return r.Response.(map[string]interface{}), nil
}
