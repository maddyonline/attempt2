package main

import (
	"log"
	"net/http"
	_ "time"
)
var mymap map[string][]byte

var state bool = false


func viewHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("State %v", state)
	state = !state
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	if state {
		w.Write(mymap["correct"])
	} else {
		w.Write(mymap["wrong"])
	}
}

func main() {
	mymap = make(map[string][]byte)
	mymap["correct"] = []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="104" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="104" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h54v20H0z"/><path fill="#4c1" d="M54 0h50v20H54z"/><path fill="url(#b)" d="M0 0h104v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">solution</text><text x="28" y="14">solution</text><text x="78" y="15" fill="#010101" fill-opacity=".3">correct</text><text x="78" y="14">correct</text></g></svg>`)
	mymap["wrong"] = []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="99" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="99" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h54v20H0z"/><path fill="#e05d44" d="M54 0h45v20H54z"/><path fill="url(#b)" d="M0 0h99v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">solution</text><text x="28" y="14">solution</text><text x="75.5" y="15" fill="#010101" fill-opacity=".3">wrong</text><text x="75.5" y="14">wrong</text></g></svg>`)

	mux := http.NewServeMux()
	mux.HandleFunc("/view", viewHandler)
	http.ListenAndServe(":8085", mux)
}
