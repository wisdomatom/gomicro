package main

import (
	"fmt"
	"context"
	"github.com/micro/go-micro/v2"
	greeter "mtest.com/micro/proto"
)

type Greeter struct {

}

func (g *Greeter)Hello(ctx context.Context, req *greeter.HelloRequest, res *greeter.HelloResponse) error  {
	res.Greeting = "Hello" + req.Name
	return nil
}

func main()  {
	service := micro.NewService(
		micro.Name("greeter"),
	)
	service.Init()

	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err:=service.Run();err!=nil{
		fmt.Println(err)
	}

}
