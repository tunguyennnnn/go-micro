package server

import (
  "fmt"
  "log"
  "net"
  "net/http"
  "net/rpc"
)

type HelloWorldRequest struct {
	Name string
}

type HelloWorldResponse struct {
	Message string
}

type HelloWorldHandler struct{}

const port = 1234

func (h *HelloWorldHandler) HelloWorld(args *HelloWorldRequest, reply *HelloWorldResponse) error {
  reply.Message = "Hello " + args.Name
	return nil
}

func StartServer() {
  helloWorld := &HelloWorldHandler{}
  rpc.Register(helloWorld)
  l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
  http.Serve(l, http.HandlerFunc(httpHandler))
	log.Printf("Server starting on port %v\n", port)

	http.Serve(l, nil)
}
