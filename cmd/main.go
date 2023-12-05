package main

import (
	"flag"
	"github.com/Markovk1n/go-kit-tutor1/pkg/handler"
	"github.com/Markovk1n/go-kit-tutor1/pkg/server"
	"github.com/Markovk1n/go-kit-tutor1/pkg/services"
	"log"
	"net/http"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)

	flag.Parse()

	srv := services.NewService()

	// создаем Endpoints
	endpoints := handler.Endpoints{
		GetEndpoint:      handler.MakeGetEndpoint(srv),
		StatusEndpoint:   handler.MakeStatusEndpoint(srv),
		ValidateEndpoint: handler.MakeValidateEndpoint(srv),
	}

	handler := server.NewHTTPServer(endpoints)

	log.Printf("dateservice is running on %s\n", *httpAddr)

	if err := http.ListenAndServe(*httpAddr, handler); err != nil {
		log.Println(err)
	}
}
