package user

import (
	"context"
	"encoding/json"

	lu "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/user"
	"github.com/quanxiang-cloud/faas-lowcode/pkg/client"
)

type user struct {
	client client.Client
}

func New(client client.Client) lu.User {
	return &user{
		client: client,
	}
}

func (u *user) GetTag() string {
	return "user"
}

func (u *user) GetHost() string {
	return "org"
}

func (u *user) GetProfile(ctx context.Context, req *lu.GetProfileReq) (*lu.GetProfileResp, error) {
	resp := new(lu.GetProfileResp)

	result, err := u.client.Do(ctx, client.WithGET(
		u,
		"/api/v1/org/m/user/info",
		map[string][]string{
			"id": {req.ID},
		},
	))

	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(result, resp)
	return resp, err
}
