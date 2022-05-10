package client

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/quanxiang-cloud/faas-lowcode/pkg/service"
	"k8s.io/klog/v2"
)

// TODO context check,user can not change header`s user id and request id
type Client interface {
	Do(ctx context.Context, opts ...Option) (Result, error)
}

var (
	timeout      time.Duration = time.Duration(getEnvToInt64WithDefatult("LOWCODE_CLIENT_TIMEOUT", 20)) * time.Second
	maxIdleConns int           = int(getEnvToInt64WithDefatult("LOWCODE_CLIENT_MAX_IDLE_CONNS", 10))
)

func New() (Client, error) {
	client := http.Client{
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
	}
	c := &httpClient{}
	c.c = client

	return c, nil
}

type client struct {
	c http.Client
}

type httpClient struct {
	client

	request *http.Request
}

func (c *httpClient) clone() *httpClient {
	client := *c
	return &client
}

func (c *httpClient) SetRequest(request *http.Request) {
	c.request = request
}

func (c *httpClient) Do(ctx context.Context, opts ...Option) (Result, error) {
	client := c.clone()

	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, err
		}
	}

	client.request.WithContext(ctx)

	response, err := client.c.Do(client.request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

type Result []byte

type Option func(Client) error

// WithGET http get method
func WithGET(svc service.Service, uri string, params interface{}) Option {
	return func(c Client) error {
		values := normalizeParams(params)

		url := url.URL{
			Scheme:   "http",
			Host:     svc.GetHost(),
			Path:     uri,
			RawQuery: values.Encode(),
		}

		request, err := http.NewRequest("GET", url.String(), nil)
		if err != nil {
			return err
		}

		c.(*httpClient).SetRequest(request)

		return nil
	}

}

func normalizeParams(params interface{}) url.Values {
	if values, ok := params.(map[string][]string); ok {
		return values
	}

	values := url.Values{}
	switch kind := reflect.TypeOf(params).Kind(); kind {
	case reflect.Struct:
		// TODO
	case reflect.Map:
		// TODO
	default:
		klog.Error("unsupported parameter type, [%v]", kind)
	}

	return values
}
