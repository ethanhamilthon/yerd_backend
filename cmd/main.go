package main

import (
	"word/config"
	"word/internal/service"
	"word/internal/storage"
	"word/internal/transport/rest"
)

func main() {
	//Load config
	config.Load()

	//Connect DB
	store, err := storage.New()
	if err != nil {
		panic("Panic: storage is not created")
	}
	defer store.CloseConnections()

	//Define services
	app := service.New(store.DB)

	//Run the server
	server := rest.New(app)
	server.Serve()
}
