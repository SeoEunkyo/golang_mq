package main

import (
	"fmt"
	"os"

	msgqueue_amqp "github.com/SeoEunkyo/golang_mq/lib/msgqueue/amqp"
	"github.com/SeoEunkyo/golang_mq/rest"
	"github.com/streadway/amqp"
)

func main (){

	amqpUrl := os.Getenv("AMQP_URL")
	if(amqpUrl == ""){
		amqpUrl = "amqp://guest:guest@3.39.245.54:5672"
	}
	conn, err := amqp.Dial(amqpUrl)
	if(err != nil){
		fmt.Println("err" + err.Error())
	}
	eventEmitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn, "events")

	rest.ServeAPI("127.0.0.1:8888", eventEmitter)
	fmt.Println("Rest API Server Start")
}