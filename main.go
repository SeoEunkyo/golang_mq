package main

import (
	"fmt"

	"github.com/SeoEunkyo/golang_mq/rest"
)

func main (){

	rest.ServeAPI("127.0.0.1:8888")
	fmt.Println("Rest API Server Start")
}