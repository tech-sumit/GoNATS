package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"nats-template/handlers"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	nats := &handlers.NATSWrapper{}
	err := nats.Start(&handlers.ClientHandler{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started successfully!")
	<-exit
	nats.Close()
	log.Println("Server stopped.")
}
