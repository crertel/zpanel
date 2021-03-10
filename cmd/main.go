package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
	"zpanel/pkg/api"

	"github.com/gorilla/mux"
)

var listenAddr string
var listenPort uint16

func init() {
	listenAddrRaw := os.Getenv("LISTEN_ADDR")
	if len(listenAddrRaw) > 0 {
		if net.ParseIP(listenAddrRaw) != nil {
			listenAddr = listenAddrRaw
		} else {
			log.Fatal(fmt.Sprintf("Bad listen address %s, exiting.", listenAddrRaw))
		}
	} else {
		listenAddr = "127.0.0.1"
	}

	listenPortRaw := os.Getenv("LISTEN_PORT")
	if len(listenPortRaw) > 0 {
		attemptedConv, err := strconv.ParseUint(listenPortRaw, 10, 16)
		if err != nil {
			log.Fatal(fmt.Sprintf("Bad listen port %s, exiting.", listenPortRaw))
		} else {
			listenPort = uint16(attemptedConv)
		}
	} else {
		listenPort = 8080
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/heartbeat", api.HeartbeatHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", listenAddr, listenPort),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
