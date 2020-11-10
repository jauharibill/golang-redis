package main

import (
	"net/http"
)

func Router() {
	handler := NewHandler(Conn())

	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/store", handler.Store)
	http.HandleFunc("/update", handler.Update)
	http.HandleFunc("/delete", handler.Delete)
	http.HandleFunc("/show", handler.Show)
}
