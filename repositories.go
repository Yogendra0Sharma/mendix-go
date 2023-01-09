package mendix

import (
	"context"
	"fmt"
	"net/http"
)

const apiurl = "https://repository.api.mendix.com/v1/repositories/"

type RepoResponse struct {
	AppID string `json:"appId"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}

func (c *Client) GetRepoInfo(
	ctx context.Context,
	appID string,
) (res RepoResponse, err error) {
	urlSuffix := fmt.Sprintf(apiurl+"%s/info", appID)
	req, err := http.NewRequest("GET", c.fullURL(urlSuffix), nil)
	if err != nil {
		return
	}

	req = req.WithContext(ctx)
	err = c.sendRequest(req, &res)
	return
}
