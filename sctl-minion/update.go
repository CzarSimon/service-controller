package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

//Service contains info about the service whose image will be updated
type Service struct {
	Image string
}

//UpdateCMD returns the command to update an image
func (service Service) UpdateCMD() string {
	return "docker pull " + service.Image
}

//UpdateImage updates the version of the supplied docker image on the minion host
func (env *Env) UpdateImage(res http.ResponseWriter, req *http.Request) {
	var service Service
	err := json.NewDecoder(req.Body).Decode(&service)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	fmt.Println(service.UpdateCMD())
	output, err := RunCommand(service.UpdateCMD())
	if err != nil {
		log.Println(output)
		util.SendErrRes(res, err)
		return
	}
	log.Println(output)
	util.SendOK(res)
}
