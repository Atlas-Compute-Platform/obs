package main

/*
	Atlas Object Storage
	Thijs Haker
*/

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/Atlas-Compute-Platform/lib"
)

var workingDirectory *string

func usage() {
	fmt.Fprintf(os.Stderr, "Atlas Object Storage %s\n", lib.VERS)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var netAddr = flag.String("p", lib.PORT, "Specify port")
	workingDirectory = flag.String("d", ".", "Specify Directory")
	flag.Usage = usage
	flag.Parse()

	if err := syscall.Chroot(*workingDirectory); err != nil {
		lib.LogError(os.Stderr, "main.main", err)
		os.Exit(1)
	}
	if err := os.Chdir("/"); err != nil {
		lib.LogError(os.Stderr, "main.main", err)
		os.Exit(1)
	}

	http.HandleFunc("/ping", lib.ApiPing)
	http.HandleFunc("/load", apiLoad)
	http.HandleFunc("/store", apiStore)
	http.HandleFunc("/remove", apiRemove)

	http.ListenAndServe(*netAddr, nil)
}
