package proxy

import (
	"context"
	"testing"
	"time"
)

func TestProxyServer(t *testing.T) {
	p, err := NewServer()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*600)
	defer cancel()

	t.Log("Starting proxy ....")
	p.Server(ctx)
	t.Log("Stop proxy")
}
