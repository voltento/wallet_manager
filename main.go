package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/voltento/pursesManager/account_managing"
	"github.com/voltento/pursesManager/database"
	"log"
	"net/http"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	r := mux.NewRouter()
	dbCtrl, err := database.CreatePsqlWalletMgr()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %v", err.Error()))
	}

	s := account_managing.CreateService(dbCtrl)
	r.Handle("/account_managing/add/", account_managing.MakeHandler(s)).Methods("PUT")

	http.Handle("/", r)
	address := ":8080"
	log.Printf("Start listen: %v", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
