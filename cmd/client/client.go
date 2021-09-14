package main

import (
	"context"
	"flag"
	"fmt"
	adrbook "github.com/aperrot42/protoreflect/adrbook/api"
	"google.golang.org/grpc"
	"net/http"
)

var serverAddr = flag.String("server address", "localhost:5999", "address and port of the gRPC server")

func main() {
	fmt.Println("client")
	flag.Parse()
	http.Post(*serverAddr)

	var opts = []grpc.DialOption{grpc.WithInsecure()}


	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := adrbook.NewAddressBookClient(conn)
	grpc.UnknownServiceHandler()
    reply, err := client.GetBook(context.Background(), &adrbook.PersonRequest{Name: "toto"})
    if err != nil {
    	panic (err)
	}
	fmt.Println(reply)
}

