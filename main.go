package main

import (
	"apimessages/src/humidity/infraestructure/adapters"
	dependenciesmessage "apimessages/src/humidity/infraestructure/dependenciesMessage"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders: []string{"Content-Type", "Authorization"},
        ExposeHeaders: []string{"Authorization"},
        MaxAge: 12 * time.Hour,
    }))

    mysqlAdapter, err := adapters.NewMySQL()
    if err != nil {
        log.Fatalf("Error al conectar con MySQL: %v", err)
    }

    webSocketAdapter := adapters.NewWebSocketAdapter()

    dependenciesmessage.InitMessages(r, webSocketAdapter, mysqlAdapter)

    if err := r.Run(":4000"); err != nil {
        panic(err)
    }
}
