package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codeedu/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandlers (r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface){
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

}	

func getProduct(service application.ProductServiceInterface) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	

}