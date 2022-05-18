package lowcode

import (
	"context"
	"fmt"
	ll "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode"
	lf "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/form"
	lu "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/quanxiang-cloud/faas-lowcode/pkg/client"
)

type LowCodeSuite struct {
	suite.Suite
	lc ll.Lowcode
	//client client.Client
	ctx   context.Context
	AppID string

	TableID string
}

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

func TestForm(t *testing.T) {
	suite.Run(t, new(LowCodeSuite))
}

func (suite *LowCodeSuite) SetupTest() {
	suite.AppID = "hrr84"
	suite.TableID = "6pd7k"

	client, err := client.New()
	assert.Nil(suite.T(), err)
	lc, err := New(client)
	assert.Nil(suite.T(), err)
	suite.lc = lc

	ctx := context.Background()
	ctx = context.WithValue(ctx, "User-Id", "f253a657-367e-4d7f-a815-94c43e327b04")
	suite.ctx = ctx
}

func (suite *LowCodeSuite) TestLowCode() {

	create, err := suite.lc.CreateForm(suite.ctx, &lf.CreateReq{
		Universal: lf.Universal{
			AppID:   suite.AppID,
			TableID: suite.TableID,
		},
		Entity: map[string]interface{}{
			"name": "name1",
			"age":  1,
		},
	})
	assert.Nil(suite.T(), err)
	fmt.Println(fmt.Sprintf("data: create is  %v ", create))
	id := getKeyToString(create.Data.Entity, "_id")

	list, err := suite.lc.ListForm(suite.ctx, &lf.ListReq{
		Universal: lf.Universal{
			AppID:   suite.AppID,
			TableID: suite.TableID,
		},
		ListOpt: lf.ListOpt{
			Page:  1,
			Size:  10,
			Query: map[string]interface{}{},
		},
	})
	assert.Nil(suite.T(), err)

	fmt.Println(fmt.Sprintf("data: list  is  %v ", list))

	get, err := suite.lc.GetForm(suite.ctx, &lf.GetReq{
		Universal: lf.Universal{
			AppID:   suite.AppID,
			TableID: suite.TableID,
		},
		ID: id,
	})
	assert.Nil(suite.T(), err)
	fmt.Println(fmt.Sprintf("data: get  is  %v ", get))

	update, err := suite.lc.UpdateForm(suite.ctx, &lf.UpdateReq{
		Universal: lf.Universal{
			AppID:   suite.AppID,
			TableID: suite.TableID,
		},
		ID: id,
		Entity: map[string]interface{}{
			"update1": "update1",
		},
	})
	assert.Nil(suite.T(), err)
	fmt.Println(fmt.Sprintf("data: update is  %v ", update))

	deletes, err := suite.lc.DeleteForm(suite.ctx, &lf.DeleteReq{
		Universal: lf.Universal{
			AppID:   suite.AppID,
			TableID: suite.TableID,
		},
		ID: id,
	})
	assert.Nil(suite.T(), err)
	fmt.Println(fmt.Sprintf("data: delete is  %v ", deletes))

}

func getKeyToString(entity lf.Entity, key string) string {
	value, ok := entity[key]
	if !ok {
		return ""
	}
	v, ok := value.(string)
	return v
}
