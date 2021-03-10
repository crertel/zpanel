package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type Heartbeat struct {
	Version     string        `json:"version"`
	Motd        string        `json:"motd"`
	Starttime   time.Time     `json:"startTime"`
	UptimeMsecs time.Duration `json:"upTimeMsecs"`
}

var version string
var motd string
var startTime time.Time

func init() {
	startTime = time.Now()
	version = "0.0.1"
	motd = "Welcome to ZPanel!"
}

func HeartbeatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&Heartbeat{
		Version:     version,
		Motd:        motd,
		Starttime:   startTime,
		UptimeMsecs: time.Since(startTime) / time.Millisecond,
	})
}
