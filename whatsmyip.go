package main

import (
	"fmt"
	"strings"
	"strconv"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	addr := r.RemoteAddr
	p := strings.LastIndex(addr, ":")
	if p != -1 {
		addr = addr[:p]
	}
	fmt.Fprint(w, addr)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "whatsmyip: expected listen port argument")
		os.Exit(1)
		return
	}

	listenPort, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "whatsmyip: listen port argument: %s\n", err.Error())
		os.Exit(1)
		return
	}

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil); err != nil {
		fmt.Fprintf(os.Stderr, "ListenAndServe: %s\n", err.Error())
		os.Exit(1)
		return
	}
}
