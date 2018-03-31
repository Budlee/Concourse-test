package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
)

const argumentPort string = "port"

type indexHandler struct{}

func (i *indexHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	type indexWelcomeJson struct {
		Message string `json:"WelcomeMessage"`
	}
	helloJson := indexWelcomeJson{
		Message: "Hello",
	}
	json.NewEncoder(writer).Encode(helloJson)

}

func startServer(port string) *http.Server {
	router := mux.NewRouter()
	router.Handle("/", &indexHandler{}).Methods("GET")

	srv := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	return &srv
}

func main() {
	log.Println("Starting", os.Args[0], "...")

	var port string
	flag.StringVar(&port, argumentPort, "8080", "The port the application should run on")
	flag.Parse()

	log.Println("Argument", argumentPort, port)

	var srv *http.Server = startServer(port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	srv.Close()
	log.Println("Shutting down")
	os.Exit(0)
}
