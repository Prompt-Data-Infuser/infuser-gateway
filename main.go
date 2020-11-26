package main

import (
	"log"
	"os"

	"gitlab.com/promptech1/infuser-gateway/client"
	"gitlab.com/promptech1/infuser-gateway/config"
	"gitlab.com/promptech1/infuser-gateway/handler"
	"gitlab.com/promptech1/infuser-gateway/router"
)

func main() {
	ballast := make([]byte, 10<<24)
	_ = ballast

	ctx := new(config.Context)
	if err := ctx.InitContext(); err != nil {
		log.Printf("Fail load config: %s", err.Error())
		os.Exit(-1)
	}

	grpcPool := client.NewGRPCPool(ctx)

	r := router.New()

	apiGroup := r.Group("/api")
	h := handler.NewHandler(grpcPool)
	h.Register(apiGroup)

	r.Logger.Fatal(r.Start(conf.Server.Host + ":" + conf.Server.Port))
}
