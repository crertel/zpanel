package api

import (
	"encoding/json"
	"net/http"

	"github.com/deniswernert/go-fstab"
)

type Fstab struct {
	Tabs fstab.Mounts `json:"filesystems"`
}

func init() {
}

func FstabHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	mounts, _ := fstab.ParseSystem()
	json.NewEncoder(w).Encode(&Fstab{
		Tabs: mounts,
	})
}
