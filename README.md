# gobee
Minimal micro service platform for gRPC and REST+Swagger APIs

Using gobee you can automatically spin up HTTP handlers for your gRPC server with complete Swagger defs with a few lines of code.

## Why?
There's an excellent [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) proxy generator, but it requires you to spin up (at least) one proxy instance in addition to your services. `gobee` allows you to serve HTTP traffic by server instances themselves for easier debugging/testing. 

It can also be used to serve production traffic, but grpc-gateway is a better fit, since you'll need an HTTP balancer for that. You can use both grpc-gateway for production and clay for testing.

## How?
First, generate your Go code using protoc:
```
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. --gobee_out=:. ./sum.proto
```
Then finish your gRPC service implementation as usual:

```
// SumImpl is an implementation of SummatorService.
type SumImpl struct{}

// Sum implements SummatorServer.Sum.
func (s *SumImpl) Sum(ctx context.Context, r *pb.SumRequest) (*pb.SumResponse, error) {
	if r.GetA() == 0 {
		return nil, errors.New("a is zero")
	}

	if r.GetB() == 65536 {
		panic(errors.New("we've got a problem!"))
	}

	sum := r.GetA() + r.GetB()
	return &pb.SumResponse{
		Sum: sum,
	}, nil
}
```

Then, add one method to the implementation, so it would implement the `"github.com/daewood/gobee/transport".Service`:
```
// GetDescription is a simple alias to the ServiceDesc constructor.
// It makes it possible to register the service implementation @ the server.
func (s *SumImpl) GetDescription() transport.ServiceDesc {
	return pb.NewSummatorServiceDesc(s)
}
```

Swagger definition will be served at `/swagger.json`.

gobee.Server is easily extendable, as you can pass any options gRPC server can use, but if it's not extendable enough then you can use the `.GetDescription()` method of your implementation to register the service in your own custom server 
