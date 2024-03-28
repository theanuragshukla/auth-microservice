package handlers

import (
	"auth-ms/middlewares"
	"encoding/json"
	"log"
	"net/http"
)

// HomeResponse is the response model for the home handler
// swagger:response HomeResponse
type HomeResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}


// swagger:route GET / home Home
// Returns a welcome message with status code 200
// responses:
// 200: HomeResponse
func (auth *Provider) HomeHandler(rw http.ResponseWriter, r *http.Request) {
	reqID := middlewares.GetTraceID(r)
	auth.L.Info(reqID)
	res := &HomeResponse{200, "Hello"}
	e := json.NewEncoder(rw)
	err := e.Encode(res)
	if err != nil {
		log.Println(err)
	}
}
