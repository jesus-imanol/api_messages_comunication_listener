package adapters

import (
	"apimessages/src/core"
	"apimessages/src/messages/domain/entities"
	"context"
	"encoding/json"
	"log"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
    conn *amqp.Connection
}

// Constructor para inyectar la conexión de RabbitMQ
func NewRabbitMQPublisher() *RabbitMQPublisher {
    conn := core.GetRabbitMQConnection()
    if conn == nil {
        log.Println("Error: RabbitMQ connection is nil")
    }
    return &RabbitMQPublisher{conn}
}

func (p *RabbitMQPublisher) InitFertilizer(ctx context.Context, message entities.MessageFertilizer) (*entities.MessageFertilizer, error) {
    if p.conn == nil {
        return nil, errors.New("RabbitMQ connection is not established")
    }
    
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
        "amq.topic",  // nombre del intercambio
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