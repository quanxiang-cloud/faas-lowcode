package main

import (
	"context"
	"fmt"
	"time"

	"github.com/quanxiang-cloud/faas-lowcode/pkg/proxy"
)

func main() {
	p, err := proxy.NewServer()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*600)
	defer cancel()

	fmt.Println("Starting ...")
	p.Server(ctx)
}
