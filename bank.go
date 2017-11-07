package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func bankInitial() int {
	errors := 0
	for i := 1; i <= 100; i++ {
		_, err := master.Insert("accounts", []interface{}{i, 100000})
		if err != nil {
			log.Println(err)
			errors++
		}
	}
	return errors
}

func problemHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := node.Call("problem", []interface{}{})
	fmt.Fprintf(w, "%v", resp)
}

func bankHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := node.Call("info", []interface{}{})
	fmt.Fprintf(w, "%v", resp)
	log.Println("Request")
}

func txHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	from, _ := strconv.Atoi(query.Get("from"))
	to, _ := strconv.Atoi(query.Get("to"))
	_, err := node.Call("send", []interface{}{from, to, 1})
	if err != nil {
		log.Printf("tx error: %v, switched to slave", err)
		node = slave
		node.Call("stop", []interface{}{})

		// heavy code emulation
		//time.Sleep(1500 * time.Millisecond)

		node.Call("send", []interface{}{from, to, 1})
	}
}
