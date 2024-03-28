// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"auth-ms/data"
	"auth-ms/handlers"
	"auth-ms/middlewares"
	"auth-ms/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	// setting up logger
	l := utils.CreateLogger()
	_ = l.Sync()

	// reading from env file
	viper.AutomaticEnv()
	viper.SetConfigName("app")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		l.Info("Error reading .env file")
	}

	db, err := data.GetDb()
	if err != nil {
		l.Info("Error connecting to db")
		l.Fatal(err.Error())
	}

	repo := utils.Repository{db}
	repo.DB.AutoMigrate(&data.User{}, &data.Session{})

	auth := handlers.NewProvider(&repo, l)

	sm := mux.NewRouter()

	sm.Use(middlewares.ReqIDMiddleware)

	getRouter := sm.Methods("GET").Subrouter()
	postRouter := sm.Methods("POST").Subrouter()

	getRouter.HandleFunc("/", auth.HomeHandler)
	getRouter.HandleFunc("/verify", auth.VerifyHandler)
	getRouter.HandleFunc("/token", auth.TokenHandler)
	getRouter.HandleFunc("/profile", auth.ProfileHandler)

	postRouter.HandleFunc("/signup", auth.SignupHandler)
	postRouter.HandleFunc("/login", auth.LoginHandler)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	docsHandler := middleware.Redoc(ops, nil)

	getRouter.Handle("/docs", docsHandler)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	s := http.Server{
		Addr:         fmt.Sprintf(":%s", viper.Get("PORT")),
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		l.Info("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Error(fmt.Sprintf("Error starting server: %s\n", err))
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	l.Info(fmt.Sprintf("Got signal: %s", sig))

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
