package Proxy

import (
	entityv1 "awesomeProject/gen/bussine"
	healthv1 "awesomeProject/gen/core"
	"awesomeProject/service"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
"google.golang.org/grpc"

"net/http"
)

func Run(s *service.Server) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var mux = runtime.NewServeMux(runtime.WithMetadata(s.Theotherfn))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := entityv1.RegisterEntityserviceAPIHandlerFromEndpoint(ctx, mux,"localhost"+s.GetPort(), opts)
	err = healthv1.RegisterHealthAPIHandlerFromEndpoint(ctx,mux,"localhost"+s.GetPort(),opts)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

