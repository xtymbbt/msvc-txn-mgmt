package main

import (
	"../proto/commonInfo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address     = "localhost:1996"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := commonInfo.NewCommonInfoClient(conn)

	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	r, err := c.SendToDataCenter(context.Background(), &commonInfo.HttpRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Recieved: %v", r.Success)

}
