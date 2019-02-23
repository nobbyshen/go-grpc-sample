# go-grpc-sample
The Go-GRPC is a basic [Go](https://github.com/golang) module with [GRPC](https://grpc.io/docs/quickstart/go.html) integration between client and server.
* Server : It hosts GRPC/HTTP server (port:8000/8001) sample code to serve client requests
* Client: It's the GRPC client sample code to test basic protobuf function
* sample.proto: It's the sample protobuf, which includes basic message, and services

# Usage
```
$ make -C client
$ make -C server
$ server/server
GRPC serve(:8000) ... 
Http serve(:8001) ...
$ client/client
$ curl http://localhost:8001/sayhello?name=hello
$ curl http://localhost:8001/SquareRoot?para1=100
```


