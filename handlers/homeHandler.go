package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HomeResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func (auth *AuthProvider) HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("request on /")
	res := &HomeResponse{200, "Hello"}
	e := json.NewEncoder(rw)
	err := e.Encode(res)
	if err != nil {
		log.Println(err)
	}
}
