// Issue 126
// Passing tainted data via map key

package main

import (
	"net/http"
	"os/exec"
)

func HandleReq(req *http.Request) {
	vals := req.URL.Query()
	cmdName := vals.Get("cmd")
	ss := map[string]string{cmdName: "command"}
	for key := range ss {
		cmd := exec.Command(key)
		cmd.Run()
	}
}
