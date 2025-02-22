package main

import (
	"context"
	"log"

	"github.com/DimTur/empty_service/cmd/serve"
)

func main() {
	ctx := context.Background()

	cmd := serve.NewServeCmd()

	if err := cmd.ExecuteContext(ctx); err != nil {
		log.Fatalf("smth went wrong: %s", err)
	}
}
