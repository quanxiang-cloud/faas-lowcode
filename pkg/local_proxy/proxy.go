package local_proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"reflect"
	"time"

	"github.com/quanxiang-cloud/faas-lowcode/pkg"
)

var (
	timeout      time.Duration = time.Duration(pkg.GetEnvToInt64WithDefatult("LOWCODE_CLIENT_TIMEOUT", 20)) * time.Second
	maxIdleConns int           = int(pkg.GetEnvToInt64WithDefatult("LOWCODE_CLIENT_MAX_IDLE_CONNS", 10))
	// host         string        = pkg.GetEnvWithDefault("LOWCODE_PROXY_HOST", "http://api.faasall.com")
	host string = pkg.GetEnvWithDefault("LOWCODE_PROXY_HOST", "http://127.0.0.1:80")
)

type proxy struct {
	client http.Client
}

func newProxy() *proxy {
	return &proxy{
		client: http.Client{
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					deadline := time.Now().Add(timeout * time.Second)
					c, err := net.DialTimeout(netw, addr, time.Second*timeout)
					if err != nil {
						return nil, err
					}
					c.SetDeadline(deadline)
					return c, nil
				},
				MaxIdleConns: maxIdleConns,
			},
		}}
}

func (p *proxy) do(funcName string, req interface{}, resp interface{}) error {
	if kind := reflect.ValueOf(resp).Kind(); kind != reflect.Ptr {
		return fmt.Errorf("expect ptr,actual %s", kind.String())
	}

	request, err := json.Marshal(req)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(request)
	respose, err := p.client.Post(fmt.Sprintf("%s/%s", host, funcName), "application/json", buf)
	if err != nil {
		return err
	}
	defer respose.Body.Close()

	if respose.StatusCode != http.StatusOK {
		return fmt.Errorf(respose.Status)
	}

	body, err := io.ReadAll(respose.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, resp)
}
