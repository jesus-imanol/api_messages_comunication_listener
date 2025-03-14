package main

import (
	"apimessages/src/message/infraestructure/adapters"
	dependenciesmessage "apimessages/src/message/infraestructure/dependenciesMessage"
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

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Authorization"}, 
		MaxAge:           12 * time.Hour,
	}))
	ps, err := adapters.NewMySQL()
	if err != nil {
	panic(err)
	}
	dependenciesmessage.InitMessages(r, ps)
	if err := r.Run(":4000"); err != nil {
		panic(err)
	}
}