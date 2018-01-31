package main

import (
	"context"

	"github.com/Sirupsen/logrus"
	example "github.com/daewood/gobee/doc/example/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	client := example.NewSummatorClient(conn)

	rsp, err := client.Sum(context.Background(), &example.SumRequest{A: 1, B: 2})
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info(rsp)
	}

	rsp, err = client.Sum(context.Background(), &example.SumRequest{A: 0, B: 2})
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info(rsp)
	}

	rsp, err = client.Sum(context.Background(), &example.SumRequest{A: 1, B: 65536})
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info(rsp)
	}

	echo := example.NewEchoServerClient(conn)
	rs, err := echo.Echo(context.Background(), &example.StringMessage{Value: "Hello world"})
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info(rs.GetValue())
	}

	// r, err := client.Wait(context.Background(), &example.WaitRequest{Name: "henry"})
	// if err != nil {
	// 	logrus.Error(err)
	// } else {
	// 	logrus.Info(r.Error)
	// }
}
