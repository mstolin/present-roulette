package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mstolin/present-roulette/gmail-adapter/handler"
	"github.com/mstolin/present-roulette/gmail-adapter/mail"
)

func main() {
	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ":8080"
	}
	listener, error := net.Listen("tcp", address)
	if error != nil {
		log.Fatalf("Error occurred: %s", error.Error())
	}

	host, port, sender, password := os.Getenv("GMAIL_HOST"), os.Getenv("GMAIL_PORT"), os.Getenv("GMAIL_MAIL"), os.Getenv("GMAIL_PASSWORD")
	smtpClient, error := mail.NewSMTPClient(host, port, sender, password)
	if error != nil {
		log.Fatalf("Could not create SMTP client: %v", error)
	}
	httpHandler := handler.NewHandler(smtpClient)
	server := &http.Server{
		Handler: httpHandler,
	}

	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Started server on %s", address)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
