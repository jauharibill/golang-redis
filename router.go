package main

import "net/http"

func Router() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/store", Store)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
}
