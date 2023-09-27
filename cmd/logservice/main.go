package main

import (
	"Distributed/miniDistributedSystem/log"
	"Distributed/miniDistributedSystem/registry"
	"Distributed/miniDistributedSystem/service"
	"context"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"

	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  fmt.Sprintf("http://%v:%v", host, port),
	}
	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
