package lowcode

import (
	"context"
	"testing"

	lu "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/user"
	"github.com/quanxiang-cloud/faas-lowcode/pkg/client"
)

func TestUser(t *testing.T) {
	client, err := client.New()
	if err != nil {
		t.Fatal(err)
	}

	lc, err := New(client)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	lc.GetProfile(ctx, &lu.GetProfileReq{})
}
