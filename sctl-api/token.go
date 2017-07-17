package main // sctl-api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	sctl "github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
	"github.com/jasonlvhit/gocron"
)

// ScheduleTokenRefresh Refreshes the API-server token at a schedueld interval
func (env *Env) ScheduleTokenRefresh(refreshFrequency uint64) {
	gocron.Every(refreshFrequency).Seconds().Do(env.RefreshToken)
	<-gocron.Start()
}

// RefreshToken Creates stores and sends a new token to all nodes
func (env *Env) RefreshToken() {
	newToken := sctl.NewToken()
	//fmt.Println("Refresh", "Old:", env.token.Data, "\nNew:", newToken.Data)
	err := env.SendToAllNodes("reset-token", &newToken)
	if err != nil {
		util.LogErr(err)
	} else {
		env.commitToken(newToken)
	}
}

// commitToken Changes the token once if no request are made with the old one
func (env *Env) commitToken(newToken sctl.Token) {
	for {
		time.Sleep(time.Second * 2)
		if env.reqCount < 1 {
			env.token = newToken
			break
		}
	}
}

// Unlock Sends and unlock request to all nodes using the project master token
func (env *Env) Unlock(res http.ResponseWriter, req *http.Request) {
	masterToken, err := getMasterToken(env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	tokens := env.token.ToBundle(masterToken)
	//tokens.Print()
	err = env.SendToAllNodes("unlock", &tokens)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}

// Lock Sends a lock request to all nodes
func (env *Env) Lock(res http.ResponseWriter, req *http.Request) {
	var tokens sctl.TokenBundle
	err := env.SendToAllNodes("lock", &tokens)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendOK(res)
}

// getMasterToken Retrives the master token of the current active project
func getMasterToken(db *sql.DB) (string, error) {
	var masterToken string
	query := "SELECT MASTER_TOKEN FROM PROJECT WHERE IS_ACTIVE=1"
	err := db.QueryRow(query).Scan(&masterToken)
	return masterToken, err
}

// GetTokens Returns the API-server token bundle to the cli
func (env *Env) GetTokens(res http.ResponseWriter, req *http.Request) {
	masterToken, err := getMasterToken(env.db)
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	jsonRes, err := json.Marshal(env.token.ToBundle(masterToken))
	if err != nil {
		util.SendErrRes(res, err)
		return
	}
	util.SendJSONRes(res, jsonRes)
}
