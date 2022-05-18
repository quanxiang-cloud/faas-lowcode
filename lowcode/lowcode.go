package lowcode

import (
	ll "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode"
	lf "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/form"
	lu "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/user"
	"github.com/quanxiang-cloud/faas-lowcode/lowcode/form"
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
	lf.Form
}

func (l *lowcode) Init(client client.Client) error {
	l.User = user.New(client)
	l.Form = form.New(client)
	return nil
}
