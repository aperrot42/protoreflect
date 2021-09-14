package main

import (
	"context"
	"flag"
	"fmt"
	adrbook "github.com/aperrot42/protoreflect/adrbook/api"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"io/ioutil"
	"math/rand"
	"net/http"
)

var addr = flag.String("url", ":5999", "listening url for the server")

func main() {
	fmt.Println("hello")
	flag.Parse()
	//server := server{}
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ccs json proto yeah")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))

		protoregistry.GlobalTypes.RangeMessages(func(t protoreflect.MessageType)bool{
			fmt.Println(t.Descriptor().FullName())
			return true})
		msgType, err := protoregistry.GlobalTypes.FindMessageByName("protoreflect.Simple")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w,err)
			return
		}
		//var msg = new(adrbook.Simple)
		msg := msgType.New().Interface()
		protojson.Unmarshal(body, msg)
		res, err := proto.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w,err)

			return
		}
		fmt.Fprint(w, res)
	})

	http.HandleFunc("/description", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ccs json proto description")
	})

	http.ListenAndServe(*addr, nil)
}



type server struct {
}

func (s server) GetBook(ctx context.Context, request *adrbook.PersonRequest) (*adrbook.PersonReply, error) {
	fmt.Printf("received request %v \n", request)
	return &adrbook.PersonReply{Person: &adrbook.Person{
		Name:        "truc",
		Id:          int32(rand.Int()),
		Email:       "somemail@toto.nc",
		Phones:      nil,
		LastUpdated: nil,
	}}, nil
}

