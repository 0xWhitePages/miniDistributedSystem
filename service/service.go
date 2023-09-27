package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

/*
集中启动服务，包括 log...
*/

// Start 启动服务
func Start(ctx context.Context, serviceName, host, port string,
	registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()

	ctx = startService(ctx, serviceName, host, port)

	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = host + ":" + port

	go func() {
		// 服务启 动
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
