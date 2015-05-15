package gollarhorn

import "net/http"

type PlatformService struct {
	client *Client
}

type PlatformResponse struct {
	Response        interface{}  `json: Response`
	ErrorCode       *int         `json: ErrorCode`
	ThrottleSeconds *float64     `json: ThrottleSeconds`
	ErrorStatus     *string      `json: ErrorStatus`
	Message         *string      `json: Message`
	MessageData     *interface{} `json: MessageData, omitempty`
}

func (ps *PlatformService) PlatformRequest(requestType, urlStr string) (*PlatformResponse, *http.Response, error) {
	req, err := ps.client.NewRequest(requestType, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(PlatformResponse)
	resp, err := ps.client.Do(req, r)

	if err != nil {
		return nil, nil, err
	}

	return r, resp, nil
}
