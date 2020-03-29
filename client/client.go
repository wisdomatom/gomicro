package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	greeter "mtest.com/micro/proto"
)

func main()  {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	g := greeter.NewGreeterService("greeter", service.Client())
	rep,err := g.Hello(context.TODO(), &greeter.HelloRequest{Name:"Atom"})
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(rep.Greeting)
}
