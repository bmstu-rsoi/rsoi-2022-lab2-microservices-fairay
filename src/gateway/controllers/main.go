package controllers

import (
	"fmt"
	"gateway/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func initControllers(r *mux.Router) {
	r.Use(utils.LogHandler)
	api1_r := r.PathPrefix("/api/v1/").Subrouter()

	InitFlights(api1_r)
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	initControllers(router)
	return router
}

func RunRouter(r *mux.Router, port uint16) error {
	c := cors.New(cors.Options{})
	handler := c.Handler(r)
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), handler)
}
