package main

import (
	"fmt"
	"log"
	"main/browser"
	"main/config"
	"net/http"
)

func main() {
	config.Load("config.yaml")
	cfg := config.Get()

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Printf("Сервер запущен на http://%s\n", addr)
	go browser.Open("http://" + addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
