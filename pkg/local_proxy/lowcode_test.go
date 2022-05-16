package local_proxy

import (
	"context"
	"testing"

	lu "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/user"
)

func TestGetProfile(t *testing.T) {
	l := New()
	lp, err := l.GetProfile(context.Background(), &lu.GetProfileReq{
		ID: "id",
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(lp)
}
