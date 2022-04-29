package main

/*
	Atlas Object Store
	Thijs Haker
*/

import (
	"flag"
	"net/http"
	"os"
	"syscall"

	"github.com/Atlas-Compute-Platform/lib"
)

var workingDirectory *string

func main() {
	lib.SvcName = "Atlas Object Store"
	lib.SvcVers = "1.0"

	var netAddr = flag.String("p", lib.PORT, "Specify port")
	workingDirectory = flag.String("d", ".", "Specify Directory")
	flag.Usage = lib.Usage
	flag.Parse()

	if err := syscall.Chroot(*workingDirectory); err != nil {
		lib.LogFatal(os.Stderr, "main.main", err)
	}
	if err := os.Chdir("/"); err != nil {
		lib.LogFatal(os.Stderr, "main.main", err)
	}

	http.HandleFunc("/ping", lib.ApiPing)
	http.HandleFunc("/load", apiLoad)
	http.HandleFunc("/store", apiStore)
	http.HandleFunc("/remove", apiRemove)

	http.ListenAndServe(*netAddr, nil)
}
