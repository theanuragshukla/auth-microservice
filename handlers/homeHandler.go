package handlers

import (
	"auth-ms/middlewares"
	"encoding/json"
	"log"
	"net/http"
)

type HomeResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func (auth *Provider) HomeHandler(rw http.ResponseWriter, r *http.Request) {
	reqID := middlewares.GetTraceID(r)
	auth.l.Info(reqID)
	res := &HomeResponse{200, "Hello"}
	e := json.NewEncoder(rw)
	err := e.Encode(res)
	if err != nil {
		log.Println(err)
	}
}
