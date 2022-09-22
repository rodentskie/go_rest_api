package server

import (
	"context"
	"flag"
	userController "go-rest-api/src/controller/users"
	db "go-rest-api/src/data"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func Server() {
	db.Start()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	// User Routes and Enpoints
	r.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userController.GetSingleUser).Methods("GET")
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
