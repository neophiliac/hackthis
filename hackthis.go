package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/yosssi/ace"
	"github.com/yosssi/ace-proxy"
)

var p = proxy.New(&ace.Options{BaseDir: "views", DynamicReload: true})
var hacks []string

func random_hack() string {
	return hacks[rand.Intn(len(hacks)-1)]
}

func handler(w http.ResponseWriter, r *http.Request) {
	tpl, err := p.Load("index", "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]string{"hack": random_hack()}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// load hacks
	content, err := ioutil.ReadFile("hacks.txt")
	if err != nil {
		log.Fatal(err)
	}
	hacks = strings.Split(string(content), "\n")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
