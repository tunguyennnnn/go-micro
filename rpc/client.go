package client

import (
  "fmt"
  "log"
  "net/rpc"
  "github.com/building-microservices-with-go/chapter1/rpc/contract"
)

const port = 1234

func CreateClient() *rpc.Client {
  client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
  if (err != nil) {
    log.Fatal("dialing:", err)
  }
  return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
  args := &contract.HelloWorldRequest{Name: "World"}
  var reply contract.HelloWorldResponse
  err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
  if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
