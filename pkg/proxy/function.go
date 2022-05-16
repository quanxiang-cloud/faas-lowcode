package proxy

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/quanxiang-cloud/faas-lowcode-interface/lowcode"
	lu "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/user"
)

type r interface{}

var funcs map[string]func(context.Context, lowcode.Lowcode, []byte) (r, error)

var fn func(context.Context, r) (r, error)

func init() {
	funcs = map[string]func(context.Context, lowcode.Lowcode, []byte) (r, error){}
	funcs["GetProfile"] = func(ctx context.Context, l lowcode.Lowcode, b []byte) (r, error) {
		req := &lu.GetProfileReq{}
		err := json.Unmarshal(b, req)
		if err != nil {
			return &lu.GetProfileResp{}, err
		}
		return l.GetProfile(ctx, req)
	}
}

func doProxy(ctx context.Context, l lowcode.Lowcode, name string, req []byte) (r, error) {
	fn, ok := funcs[name]
	if !ok {
		return nil, fmt.Errorf("unknown function [%s]", name)
	}

	return fn(ctx, l, req)
}
