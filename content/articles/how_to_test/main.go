package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

var counter = map[string]int{}
var mu sync.Mutex // mutex for counter

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	mu.Lock()
	counter[name]++
	cnt := counter[name]
	mu.Unlock()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome!</h1> <p>Name: " + name + "</p> <p>Count: " + fmt.Sprint(cnt) + "</p>"))

	logrus.WithFields(logrus.Fields{
		"module": "main",
		"name":   name,
		"count":  cnt,
	}).Infof("visited")
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	http.HandleFunc("/hello", handleHello)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
