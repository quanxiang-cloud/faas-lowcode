package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/quanxiang-cloud/faas-lowcode/pkg"
	"github.com/quanxiang-cloud/faas-lowcode/pkg/service"
	"k8s.io/klog/v2"
)

// TODO context check,user can not change header`s user id and request id
type Client interface {
	Do(ctx context.Context, opts ...Option) (Result, error)
}

const (
	safe   = 0
	unsafe = 1
)

var (
	timeout      time.Duration = time.Duration(pkg.GetEnvToInt64WithDefatult("LOWCODE_CLIENT_TIMEOUT", 20)) * time.Second
	maxIdleConns int           = int(pkg.GetEnvToInt64WithDefatult("LOWCODE_CLIENT_MAX_IDLE_CONNS", 10))
	hostSuffix   string        = pkg.GetEnv("LOWCODE_NAMESPACE")
	safeUser     int64         = pkg.GetEnvToInt64WithDefatult("LOWCODE_USER_SAFE", safe)
)

var (
	ErrChangeUserForbidden = errors.New("Illegal modification of user attributes")
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
	if safeUser == safe && !isAllowedThrough(ctx) {
		return nil, ErrChangeUserForbidden
	}

	client := c.clone()

	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, err
		}
	}

	client.request = client.request.WithContext(ctx)

	setHeader(client.request, ctx)

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
		host := svc.GetHost()
		if hostSuffix != "" {
			host += hostSuffix
		}
		url := url.URL{
			Scheme: "http",
			Host:   host,
			Path:   uri,
		}
		if values != nil {
			url.RawQuery = values.Encode()
		}
		request, err := http.NewRequest(http.MethodGet, url.String(), nil)
		if err != nil {
			return err
		}

		c.(*httpClient).SetRequest(request)

		return nil
	}
}

// WithDelete http delete method
func WithDelete(svc service.Service, uri string, params interface{}) Option {
	return func(c Client) error {
		values := normalizeParams(params)
		host := svc.GetHost()
		if hostSuffix != "" {
			host += hostSuffix
		}
		url := url.URL{
			Scheme: "http",
			Host:   host,
			Path:   uri,
		}
		if values != nil {
			url.RawQuery = values.Encode()
		}
		request, err := http.NewRequest(http.MethodDelete, url.String(), nil)
		if err != nil {
			return err
		}
		c.(*httpClient).SetRequest(request)
		return nil
	}
}

// WithPut http put method
func WithPut(svc service.Service, uri string, params interface{}) Option {
	return func(c Client) error {
		reader, err := bodyParams(params)
		if err != nil {
			return err
		}
		host := svc.GetHost()
		if hostSuffix != "" {
			host += hostSuffix
		}
		url := url.URL{
			Scheme: "http",
			Host:   host,
			Path:   uri,
		}
		request, err := http.NewRequest(http.MethodPut, url.String(), reader)
		if err != nil {
			return err
		}
		c.(*httpClient).SetRequest(request)
		return nil
	}
}

// WithPost http put method
func WithPost(svc service.Service, uri string, params interface{}) Option {
	return func(c Client) error {
		reader, err := bodyParams(params)
		if err != nil {
			return err
		}
		host := svc.GetHost()
		if hostSuffix != "" {
			host += hostSuffix
		}
		url := url.URL{
			Scheme: "http",
			Host:   host,
			Path:   uri,
		}
		request, err := http.NewRequest(http.MethodPost, url.String(), reader)
		if err != nil {
			return err
		}
		c.(*httpClient).SetRequest(request)
		return nil
	}
}

func normalizeParams(params interface{}) url.Values {
	if params == nil {
		return nil
	}
	if values, ok := params.(map[string][]string); ok {
		return values
	}

	values := url.Values{}
	switch kind := reflect.TypeOf(params).Kind(); kind {
	// TODO
	case reflect.Struct:
	// TODO
	case reflect.Map:

	default:
		klog.Error("unsupported parameter type, [%v]", kind)
	}

	return values
}

func bodyParams(params interface{}) (*bytes.Reader, error) {
	paramByte, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(paramByte), nil
}

func setHeader(req *http.Request, ctx context.Context) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(GetRequestIDKV(ctx).Wreck())
	req.Header.Add(GetTimezone(ctx).Wreck())
	req.Header.Add(GetTenantID(ctx).Wreck())
	req.Header.Add(GetUserID(ctx).Wreck())
}
