package local_proxy

import (
	"context"
	ll "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode"
	lf "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/form"
	lu "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/user"
)

func New() ll.Lowcode {
	lc := new(lowcode)

	lc.p = *newProxy()
	return lc
}

type lowcode struct {
	p proxy
}

func (l *lowcode) ListForm(ctx context.Context, req *lf.ListReq) (*lf.ListResp, error) {
	resp := new(lf.ListResp)
	err := l.p.do("ListForm", req, resp)
	if err != nil {
		return &lf.ListResp{}, err
	}
	return resp, nil
}

func (l *lowcode) GetForm(ctx context.Context, req *lf.GetReq) (*lf.GetResp, error) {
	resp := new(lf.GetResp)
	err := l.p.do("ListForm", req, resp)
	if err != nil {
		return &lf.GetResp{}, err
	}
	return resp, nil
}

func (l *lowcode) DeleteForm(ctx context.Context, req *lf.DeleteReq) (*lf.DeleteResp, error) {
	resp := new(lf.DeleteResp)
	err := l.p.do("DeleteForm", req, resp)
	if err != nil {
		return &lf.DeleteResp{}, err
	}
	return resp, nil
}

func (l *lowcode) UpdateForm(ctx context.Context, req *lf.UpdateReq) (*lf.UpdateResp, error) {
	resp := new(lf.UpdateResp)
	err := l.p.do("UpdateForm", req, resp)
	if err != nil {
		return &lf.UpdateResp{}, err
	}
	return resp, nil
}

func (l *lowcode) CreateForm(ctx context.Context, req *lf.CreateReq) (*lf.CreateRsp, error) {
	resp := new(lf.CreateRsp)
	err := l.p.do("CreateForm", req, resp)
	if err != nil {
		return &lf.CreateRsp{}, err
	}
	return resp, nil
}

func (l *lowcode) GetProfile(ctx context.Context, req *lu.GetProfileReq) (*lu.GetProfileResp, error) {
	resp := new(lu.GetProfileResp)

	err := l.p.do("GetProfile", req, resp)
	if err != nil {
		return &lu.GetProfileResp{}, err
	}
	return resp, nil
}
