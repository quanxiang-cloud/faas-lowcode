package lowcode

import (
	ll "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode"
	lu "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/user"
	"github.com/quanxiang-cloud/faas-lowcode/lowcode/user"
	"github.com/quanxiang-cloud/faas-lowcode/pkg/client"
)

func New(client client.Client) (ll.Lowcode, error) {
	lc := new(lowcode)
	err := lc.Init(client)
	if err != nil {
		return nil, err
	}

	return lc, nil
}

type lowcode struct {
	lu.User
}

func (l *lowcode) Init(client client.Client) error {
	l.User = user.New(client)
	return nil
}
