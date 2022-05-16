package plugin_quanxiang_lowcode_client

import (
	"context"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
	"github.com/OpenFunction/functions-framework-go/plugin"
	"github.com/fatih/structs"
	ll "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode"
	lowcode "github.com/quanxiang-cloud/faas-lowcode/pkg/local_proxy"
	"k8s.io/klog/v2"
)

const (
	Name    = "plugin-quanxiang-lowcode-client"
	Version = "v1"
)

type PluginLowcodeClient struct {
	lcode ll.Lowcode
}

var _ plugin.Plugin = &PluginLowcodeClient{}

var defaultClient *PluginLowcodeClient

func init() {
	defaultClient = &PluginLowcodeClient{}

	defaultClient.lcode = lowcode.New()
	klog.Info("init lowcode success")

}

func New() *PluginLowcodeClient {
	return defaultClient
}

func (p *PluginLowcodeClient) Name() string {
	return Name
}

func (p *PluginLowcodeClient) Version() string {
	return Version
}

func (p *PluginLowcodeClient) Init() plugin.Plugin {
	return New()
}

// TODO context check,user can not change header`s user id and request id
func (p *PluginLowcodeClient) ExecPreHook(ofCtx ofctx.RuntimeContext, plugins map[string]plugin.Plugin) error {
	request := ofCtx.GetSyncRequest().Request

	ctx := context.WithValue(request.Context(), ll.LOWCODE, p.lcode)
	r := request.WithContext(ctx)
	ofCtx.SetSyncRequest(ofCtx.GetSyncRequest().ResponseWriter, r)
	return nil
}

func (p *PluginLowcodeClient) ExecPostHook(ctx ofctx.RuntimeContext, plugins map[string]plugin.Plugin) error {
	return nil
}

func (p *PluginLowcodeClient) Get(fieldName string) (interface{}, bool) {
	plgMap := structs.Map(p)
	value, ok := plgMap[fieldName]
	return value, ok
}
