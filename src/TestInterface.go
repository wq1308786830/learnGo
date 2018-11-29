package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Counter int

var addr = flag.String("addr", ":5757", "http service address")
var counter = new(Counter)

// 简单的计数器服务。
func (counter *Counter) PrintCount(w http.ResponseWriter) {
	*counter++
	fmt.Fprintf(w, "counter = %d\n", *counter)
}

// 实参服务器。
func ArgServer(w http.ResponseWriter, req *http.Request) {
	counter.PrintCount(w)
	fmt.Printf("\n%s %s %s %s\n", req.Method, req.Proto, req.RequestURI, req.Body)
}

func main() {
	http.Handle("/", http.HandlerFunc(ArgServer))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
