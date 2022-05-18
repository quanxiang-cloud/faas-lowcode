package plugin_quanxiang_lowcode_client

import (
	"context"
	"net/http"
	"time"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
	"github.com/OpenFunction/functions-framework-go/plugin"
	"github.com/fatih/structs"
	ll "github.com/quanxiang-cloud/faas-lowcode-interface/lowcode"
	lowcode "github.com/quanxiang-cloud/faas-lowcode/lowcode"
	lc "github.com/quanxiang-cloud/faas-lowcode/pkg/client"
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

	lclient, err := lc.New()
	if err != nil {
		klog.Errorf("init service client,Err: %v", err)
	}

	defaultClient.lcode, err = lowcode.New(lclient)
	if err != nil {
		klog.Errorf("init lowcode client,Err: %v", err)
	}

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

	ctx := p.mutateContext(request.Context(), request)

	r := request.WithContext(ctx)
	ofCtx.SetSyncRequest(ofCtx.GetSyncRequest().ResponseWriter, r)

	klog.InfoS(r.RequestURI, "method", r.Method, "time", time.Now().Format("2006-01-02T15:04:05.999Z"), "requestID", ctx.Value("Request-Id"))
	return nil
}

func (p *PluginLowcodeClient) mutateContext(ctx context.Context, request *http.Request) context.Context {
	ctx = context.WithValue(ctx, ll.LOWCODE, p.lcode)

	// quanxiang lowcode context, necessary ID pass.
	var (
		_requestID interface{} = "Request-Id"
		_timezone  interface{} = "Timezone"
		_userID    interface{} = "User-Id"
	)

	ctx = context.WithValue(ctx, _requestID, request.Header.Get("Request-Id"))
	ctx = context.WithValue(ctx, _timezone, request.Header.Get("Timezone"))
	ctx = context.WithValue(ctx, _userID, request.Header.Get("User-Id"))

	ctx = lc.WithFingerprint(ctx)

	return ctx
}

func (p *PluginLowcodeClient) ExecPostHook(ofctx ofctx.RuntimeContext, plugins map[string]plugin.Plugin) error {
	ctx := ofctx.GetContext().Ctx
	klog.InfoS("finish", "time", time.Now().Format("2006-01-02T15:04:05.999Z"), "requestID", ctx.Value("Request-Id"))
	return nil
}

func (p *PluginLowcodeClient) Get(fieldName string) (interface{}, bool) {
	plgMap := structs.Map(p)
	value, ok := plgMap[fieldName]
	return value, ok
}
