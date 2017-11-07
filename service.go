package main

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
	"time"
)

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\": \"OK\"}")
}

const (
	masterUri string = "127.0.0.1:3315"
	slaveUri  string = "127.0.0.1:3316"
)

var (
	node   *tarantool.Connection
	master *tarantool.Connection
	slave  *tarantool.Connection
	err    error
)

func connect(uri string) *tarantool.Connection {
	var conn *tarantool.Connection
	opts := tarantool.Opts{
		Timeout: 1000 * time.Millisecond,
	}

	maxTry := 10
	i := 0

	for i < maxTry {
		conn, err = tarantool.Connect(uri, opts)
		if err != nil {
			log.Printf("Node '%s' is unavailable\n", uri)
			time.Sleep(1000 * time.Millisecond)
		} else {
			log.Printf("Connected to %s\n", uri)
			break
		}
		i++
	}
	return conn
}

func main() {
	master = connect(masterUri)
	slave = connect(slaveUri)

	count := bankInitial()

	log.Printf("Load errors: %d\n", count)
	node = master

	http.HandleFunc("/", handler)
	http.HandleFunc("/bank", bankHandler)
	http.HandleFunc("/tx", txHandler)
	http.HandleFunc("/problem", problemHandler)
	http.HandleFunc("/_info", info)
	log.Println("Ready 8890")
	err = http.ListenAndServe(":8890", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
