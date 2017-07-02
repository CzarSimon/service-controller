package main

import (
	"fmt"

	"github.com/CzarSimon/util"
)

func sendToMinion(minion util.ServerConfig, route string, jsonBody []byte) {
	endpoint := minion.ToURL(route)
	fmt.Println(endpoint)
	fmt.Println(string(jsonBody))
}
