package adapters

import (
	"apimessages/src/core"
	"apimessages/src/messages/domain/entities"
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
    conn *amqp.Connection
}

// Constructor para inyectar la conexión de RabbitMQ
func NewRabbitMQPublisher() *RabbitMQPublisher {
    conn := core.GetRabbitMQConnection()
    return &RabbitMQPublisher{conn}
}

func (p *RabbitMQPublisher) InitFertilizer(ctx context.Context, message entities.MessageFertilizer) ( *entities.MessageFertilizer, error) {
    ch, err := p.conn.Channel()
    if err != nil {
        return nil, err
    }
    defer ch.Close()

    body, err := json.Marshal(message)
    if err != nil {
        return nil, err
    }


    err = ch.PublishWithContext(ctx,
        "amqp.topic",  // nombre del intercambio
        "irrigation.control",  // clave de enrutamiento
        false,         // mandatory
        false,         // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
    if err != nil {
        return nil, err
    }

    log.Printf("[x] Sent: %s", body)
    return &message, nil
}