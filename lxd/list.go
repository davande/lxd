package main

import (
	"net/http"

	"gopkg.in/lxc/go-lxc.v2"

	"github.com/lxc/lxd"
)

func listGet(d *Daemon, r *http.Request) Response {
	lxd.Debugf("responding to list")

	result := make([]string, 0)

	containers := lxc.DefinedContainers(d.lxcpath)
	for i := range containers {
		result = append(result, containers[i].Name())
	}

	return SyncResponse(true, result)
}

var listCmd = Command{"list", false, false, listGet, nil, nil, nil}
