package main

import (
	"context"
	"github.com/basterrus/go_backend_framework/internal/app"
	"github.com/basterrus/go_backend_framework/internal/config"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.GetConfig()

	application, err := app.NewApplication(ctx, cfg)
	if err != nil {
		panic(err)
	}

	err = application.Run()
	if err != nil {
		panic(err)
	}
}
