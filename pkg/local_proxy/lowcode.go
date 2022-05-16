package local_proxy

import (
	"context"

	ll "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode"
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

func (l *lowcode) GetProfile(ctx context.Context, req *lu.GetProfileReq) (*lu.GetProfileResp, error) {
	resp := new(lu.GetProfileResp)

	err := l.p.do("GetProfile", req, resp)
	if err != nil {
		return &lu.GetProfileResp{}, err
	}
	return resp, nil
}
