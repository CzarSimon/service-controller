package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

func sendToMinion(minion util.ServerConfig, route string, jsonBody []byte) {
	endpoint := minion.ToURL(route)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		util.LogErr(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		util.LogErr(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return
	}
	log.Println("Non 200 response")
}
