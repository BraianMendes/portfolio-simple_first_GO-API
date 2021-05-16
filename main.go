package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BraianMendes/FirstGoAPI/house"
)

func main() {
	s := house.NewInMemoryStorage()
	router := http.NewServeMux()

	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"service is up and running"}`))
		return
	})

	router.HandleFunc("/lightbulbs", house.GetLightbulb(s))
	router.HandleFunc("/lightbulbs/create", house.CreateLightbulb(s))
	router.HandleFunc("/lightbulbs/switch", house.SwitchLightbulb(s))
	router.HandleFunc("/lightbulbs/delete", house.DeleteLightbulb(s))

	srv := http.Server{
		Addr:         ":8080",
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
		Handler:      router,
	}

	fmt.Println("http server listening on localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
