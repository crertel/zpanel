package api

import (
	"encoding/json"
	"net/http"

	"github.com/shirou/gopsutil/v3/process"
)

type Process struct {
	Pid        int     `json:"pid"`
	Cmd        string  `json:"cmd"`
	Name       string  `json:"name"`
	PPid       int     `json:"parentPid"`
	Status     string  `json:"status"`
	Priority   int32   `json:"priority"`
	NumThreads int32   `json:"numThreads"`
	NumFDs     int32   `json:"numFDs"`
	MemPercent float32 `json:"memPercent"`
}

type Processes struct {
	Procs []Process `json:"processes"`
}

func init() {
}

func rehydrateProcs(procs []*process.Process) []Process {
	ret := make([]Process, len(procs))
	for i, proc := range procs {
		cmd, _ := proc.Cmdline()
		ppid, _ := proc.Ppid()
		status, _ := proc.Status()
		nice, _ := proc.Nice()
		numThreads, _ := proc.NumThreads()
		numFds, _ := proc.NumFDs()
		memPercent, _ := proc.MemoryPercent()
		name, _ := proc.Name()
		ret[i] = Process{
			Pid:        int(proc.Pid),
			Cmd:        cmd,
			PPid:       int(ppid),
			Status:     status[0],
			Priority:   nice,
			NumThreads: numThreads,
			NumFDs:     numFds,
			MemPercent: memPercent,
			Name:       name,
		}
	}
	return ret
}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	procs, _ := process.Processes()

	json.NewEncoder(w).Encode(&Processes{
		Procs: rehydrateProcs(procs),
	})
}
