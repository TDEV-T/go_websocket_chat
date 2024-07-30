package main

import (
	"context"
	"log"
	"net/http"
)

func main() {

	Rctx := context.Background()
	ctx, cancel := context.WithCancel(Rctx)

	defer cancel()

	setupAPI(ctx)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Err: ", err)
	}
}

func setupAPI(ctx context.Context) {

	manager := NewManager(ctx)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	http.HandleFunc("/ws", manager.serverWS)
	http.HandleFunc("/login", manager.loginHandler)
}
