package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codeedu/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)


type Webserver struct{
	Service application.ProductServiceInterface
}

func MakeNewWebServer()*Webserver{
	return &Webserver{}
}

	r := mux.NewRouter() 
	n := negroni.New(
		negroni.NewLogger(),

	)

func (w Webserver) Serve(){
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 & time.Second,
		Addr: ":8080",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log:", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}