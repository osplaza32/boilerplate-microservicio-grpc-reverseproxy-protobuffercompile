package main

import (
	"awesomeProject/Proxy"
	entityv1 "awesomeProject/gen/bussine"
	healthv1 "awesomeProject/gen/core"
	"awesomeProject/servicegrpc"
	"google.golang.org/grpc/reflection"

	"fmt"
)

func main() {
	ser,err := servicegrpc.NewServer()
	if err != nil {
		fmt.Println(err)
	}
	grsp,let,erro:=ser.Start()
	if erro != nil {
		fmt.Println(err)
	}
	go Proxy.Run(ser)

	entityv1.RegisterEntityserviceAPIServer(grsp,ser)
	healthv1.RegisterHealthAPIServer(grsp,ser)
	reflection.Register(grsp)
	ser.GetLogUber().Info("GET STARED SERVER")

	if err := grsp.Serve(let); err != nil {

	}
}