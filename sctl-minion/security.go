package main // sctl-minion

import (
	"fmt"
	"net/http"
)

// ValidToken Checks if request token equals the minion token
func (env Env) ValidToken(req *http.Request) bool {
	reqToken := req.Header.Get("Authorization")
	fmt.Println("minon token:", env.token)
	fmt.Println("Request token:", reqToken)
	return env.token == reqToken
}
