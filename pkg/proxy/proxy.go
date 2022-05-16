package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	lowcode "github.com/quanxiang-cloud/faas-lowcode/lowcode"
	lc "github.com/quanxiang-cloud/faas-lowcode/pkg/client"
)

var (
	port = 80
)

type Proxy struct {
	server *http.Server
}

func NewServer() (*Proxy, error) {
	mux := http.NewServeMux()

	p := &Proxy{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}

	err := p.setHandle(mux)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Proxy) setHandle(mux *http.ServeMux) error {
	lClient, err := lc.New()
	if err != nil {
		return err
	}

	l, err := lowcode.New(lClient)
	if err != nil {
		return err
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Content-Type"), "json") {
			w.Write([]byte(`Content-Type must is "application/json"`))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			// w.Write([]byte(fmt.Sprintf("read body err: %s", err.Error())))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		name := strings.TrimPrefix(r.URL.Path, "/")
		result, err := doProxy(ctx, l, name, body)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("do proxy err: %s, name: %s", err.Error(), name)))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		response, err := json.Marshal(result)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("json marshal result err: %s, name: %s", err.Error(), name)))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Write(response)
	})
	return nil
}

func (p *Proxy) Server(ctx context.Context) {
	doneCh := make(chan struct{})

	go func() {
		select {
		case <-ctx.Done():
			shutdownCtx, cancel := context.WithTimeout(
				context.Background(),
				time.Second*5,
			)
			defer cancel()
			p.server.Shutdown(shutdownCtx)
		case <-doneCh:
		}
	}()

	err := p.server.ListenAndServe()
	if err != http.ErrServerClosed {
		panic(err)
	}
	close(doneCh)
}
