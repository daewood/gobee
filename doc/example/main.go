package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/pressly/chi"
	"github.com/rakyll/statik/fs"

	pb "github.com/daewood/gobee/doc/example/pb"
	"github.com/daewood/gobee/log"
	"github.com/daewood/gobee/transport"
	"github.com/daewood/gobee/transport/middlewares/mwgrpc"
	"github.com/daewood/gobee/transport/middlewares/mwhttp"
	"github.com/daewood/gobee/transport/server"
	"golang.org/x/net/context"

	// We're using statik-compiled files of Swagger UI
	// for the sake of example.
	_ "github.com/daewood/gobee/static/statik"
)

// SumImpl is an implementation of SummatorService.
type SumImpl struct{}

// Sum implements SummatorServer.Sum.
func (s *SumImpl) Sum(ctx context.Context, r *pb.SumRequest) (*pb.SumResponse, error) {
	if r.GetA() == 0 {
		return nil, errors.New("a is zero")
	}

	if r.GetB() == 65536 {
		panic(errors.New("we've got a problem"))
	}

	sum := r.GetA() + r.GetB()
	return &pb.SumResponse{
		Sum: sum,
	}, nil
}

// Wait implements SummatorServer.Wait.
func (s *SumImpl) Wait(ctx context.Context, r *pb.WaitRequest) (*pb.WaitResponse, error) {
	fmt.Println(r.GetName(), " Start to sleep 100 sec")
	time.Sleep(100 * time.Second)
	fmt.Println(r.GetName(), " Stop sleeping")

	return &pb.WaitResponse{
		Error: "OK",
	}, nil
}

// Hello implements SummatorServer.Hello.
func (s *SumImpl) Hello(ctx context.Context, r *pb.StrMessage) (*pb.StrMessage, error) {
	return &pb.StrMessage{
		Value: "Hello:" + r.Value,
	}, nil
}

// GetDescription is a simple alias to the ServiceDesc constructor.
// It makes it possible to register the service implementation @ the server.
func (s *SumImpl) GetDescription() transport.ServiceDesc {
	return pb.NewSummatorServiceDesc(s)
}

// EchoImpl is an implementation of EchoService.
type EchoImpl struct{}

// Echo implements EchoServer.Echo
func (s *EchoImpl) Echo(ctx context.Context, r *pb.StringMessage) (*pb.StringMessage, error) {
	return r, nil
}

// GetDescription is a simple alias to the ServiceDesc constructor.
func (s *EchoImpl) GetDescription() transport.ServiceDesc {
	return pb.NewEchoServerServiceDesc(s)
}

func main() {
	// Wire up our bundled Swagger UI
	staticFS, err := fs.New()
	if err != nil {
		logrus.Fatal(err)
	}
	hmux := chi.NewRouter()
	hmux.Mount("/", http.FileServer(staticFS))

	sumimpl := &SumImpl{}
	srv := server.NewServer(
		8800,
		// Need different http port
		//server.WithHTTPPort(9000),
		// Pass our mux with Swagger UI
		server.WithHTTPMux(hmux),
		// Recover from HTTP panics
		server.WithHTTPMiddlewares(mwhttp.Recover(log.Default), mwhttp.CloseNotifier()),
		// Recover from gRPC panics
		server.WithGRPCUnaryMiddlewares(mwgrpc.UnaryPanicHandler(log.Default)),
	)

	echoimpl := &EchoImpl{}
	srvImpls := []transport.Service{sumimpl, echoimpl}

	//Single service
	//err = srv.Run(sumimpl)

	//multi services
	err = srv.Runs(srvImpls)
	if err != nil {
		logrus.Fatal(err)
	}
}
