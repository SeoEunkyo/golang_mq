package amqp

import (
	"encoding/json"
	"fmt"

	"github.com/SeoEunkyo/golang_mq/lib/msgqueue"
	"github.com/streadway/amqp"
)

type amqpEventEmitter struct {
	connection *amqp.Connection
	exchange   string
	events     chan *emittedEvent
}
type emittedEvent struct {
	event     msgqueue.Event
	errorChan chan error
}

func NewAMQPEventEmitter(conn *amqp.Connection, exchange string)(msgqueue.EventEmitter, error){
	emitter := amqpEventEmitter{
		connection: conn,
		exchange:   exchange,
	}
	err := emitter.setup()
	if err != nil {
		return nil, err
	}

	return &emitter, nil	
}

func (a *amqpEventEmitter) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	// Normally, all(many) of these options should be configurable.
	// For our example, it'll probably do.
	err = channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil)
	return err
}

func (a *amqpEventEmitter) Emit(event msgqueue.Event) error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	// TODO: Alternatives to JSON? Msgpack or Protobuf, maybe?
	jsonBody, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("could not JSON-serialize event: %s", err)
	}

	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		ContentType: "application/json",
		Body:        jsonBody,
	}

	err = channel.Publish(a.exchange, event.EventName(), false, false, msg)
	return err
}