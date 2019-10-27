package service

import (
	entityv1 "awesomeProject/gen/bussine"
	healthv1 "awesomeProject/gen/core"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"net/http"
)


type Server struct {
	addr       string
	loguber    *zap.Logger
}

func (s *Server) Check(context.Context, *healthv1.CheckRequest) (*healthv1.CheckResponse, error) {
	panic("implement me")
}

func (s *Server) Watch(*healthv1.CheckRequest, healthv1.HealthAPI_WatchServer) error {
	panic("implement me")
}

func (s *Server)Theotherfn(ctx context.Context,req *http.Request) metadata.MD{
	as := metadata.MD{}
	as.Append("Hola","CHAO")
	return as
}

func (s *Server) Entity(ctx context.Context, req *entityv1.EntityRequest) (*entityv1.EntityResponse, error) {
	panic("implement me")
}

func (s *Server) GetEntity(ctx context.Context,req *entityv1.GetEntityRequest) (*entityv1.GetEntityResponse, error) {
	return &entityv1.GetEntityResponse{MyBool:true,Code:"200"},nil

}

type ServerOption func(*Server) error

func NewServer(opts ...ServerOption) (*Server, error) {
	logger, _ := zap.NewProduction()

	s := &Server{
		addr: ":50051",
		loguber:logger,
	}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func SetAddr(addr string) ServerOption {
	return func(s *Server) error {
		s.addr = addr
		return nil
	}
}

func (s *Server) Start() (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return nil, nil, err
	}
	gs := grpc.NewServer()
	return gs, lis,nil
}

func (s *Server) GetPort() string{
	return s.addr
}
func (s *Server) GetLogUber() *zap.Logger {
	return s.loguber
}
