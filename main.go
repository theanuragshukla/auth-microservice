package main

import (
	"auth-ms/data"
	"auth-ms/handlers"
	"auth-ms/utils"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	l := log.New(os.Stdout, "microservice", log.LstdFlags)

	db, err := data.GetDb()
	if err != nil {
		l.Fatal(err)
		return
	}

	repo := utils.Repository{db}
	repo.DB.AutoMigrate(&data.User{}, &data.Session{})

	auth := handlers.NewAuthProvider(&repo, l)

	sm := mux.NewRouter()
	getRouter := sm.Methods("GET").Subrouter()
	postRouter := sm.Methods("POST").Subrouter()

	getRouter.HandleFunc("/", auth.HomeHandler)
	getRouter.HandleFunc("/verify", auth.VerifyHandler)
	getRouter.HandleFunc("/token", auth.TokenHandler)

	postRouter.HandleFunc("/signup", auth.SignupHandler)
	postRouter.HandleFunc("/login", auth.LoginHandler)

	s := http.Server{
		Addr:         fmt.Sprintf(":%s", viper.Get("PORT")),
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
