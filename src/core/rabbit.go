package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitMQConn *amqp.Connection

func InitRabbitMQConnection() {
	errLoad := godotenv.Load()

	if errLoad != nil {
		log.Fatal("Error loading.env file")
	}

	rmqHost := os.Getenv("RMQ_HOST")
	rmqUser := os.Getenv("RMQ_USER")
	rmqPass := os.Getenv("RMQ_PASS")
	rmqPort := os.Getenv("RMQ_PORT")
	if rmqHost == "" || rmqUser == "" || rmqPass == "" || rmqPort == "" {
		log.Fatal("Faltan una o más variables de entorno necesarias para la conexión a RabbitMQ")
	}
	rmqURL := "amqp://" + rmqUser + ":" + rmqPass + "@" + rmqHost + ":" + rmqPort + "/"

	// Establecer la conexión a RabbitMQ
	var err error
	rabbitMQConn, err = amqp.Dial(rmqURL)
	if err != nil {
		log.Fatalf("Fallo al conectarse a RabbitMQ: %s", err)
	}
}

func GetRabbitMQConnection() *amqp.Connection {
	return rabbitMQConn
}
