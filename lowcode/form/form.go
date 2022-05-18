package form

import (
	"context"
	"encoding/json"
	"fmt"

	lf "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/form"
	"github.com/quanxiang-cloud/faas-lowcode/pkg/client"
)

type form struct {
	client client.Client
}

func (f *form) GetForm(ctx context.Context, req *lf.GetReq) (*lf.GetResp, error) {
	resp := new(lf.GetResp)

	result, err := f.client.Do(ctx, client.WithGET(
		f,
		getURL(req.AppID, req.TableID, req.ID),
		nil,
	))
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(result, resp)
	return resp, err
}

func (f *form) DeleteForm(ctx context.Context, req *lf.DeleteReq) (*lf.DeleteResp, error) {
	resp := new(lf.DeleteResp)

	result, err := f.client.Do(ctx, client.WithDelete(
		f,
		getURL(req.AppID, req.TableID, req.ID),
		nil,
	))
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(result, resp)
	return resp, err
}

func (f *form) UpdateForm(ctx context.Context, req *lf.UpdateReq) (*lf.UpdateResp, error) {
	resp := new(lf.UpdateResp)

	result, err := f.client.Do(ctx, client.WithPut(
		f,
		getURL(req.AppID, req.TableID, req.ID),
		map[string]interface{}{"entity": req.Entity},
	))
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(result, resp)
	return resp, err
}

func (f *form) CreateForm(ctx context.Context, req *lf.CreateReq) (*lf.CreateRsp, error) {
	resp := new(lf.CreateRsp)

	result, err := f.client.Do(ctx, client.WithPost(
		f,
		getURLs(req.AppID, req.TableID),
		map[string]interface{}{"entity": req.Entity},
	))
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(result, resp)
	return resp, err
}

func New(client client.Client) lf.Form {
	return &form{
		client: client,
	}
}

func (f *form) GetTag() string {
	return "form"
}

func (f *form) GetHost() string {
	return "form"
}

func (f *form) ListForm(ctx context.Context, req *lf.ListReq) (*lf.ListResp, error) {
	resp := new(lf.ListResp)

	result, err := f.client.Do(ctx, client.WithGET(
		f,
		getURLs(req.AppID, req.TableID),
		nil,
	))
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(result, resp)
	return resp, err
}

func getURL(appID, tableID, id string) string {
	return fmt.Sprintf("/api/v2/form/%s/home/form/%s/%s", appID, tableID, id)
}
func getURLs(appID, tableID string) string {
	return fmt.Sprintf("/api/v2/form/%s/home/form/%s", appID, tableID)
}
