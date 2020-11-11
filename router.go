package main

import (
	"github.com/vmihailenco/treemux"
)

func Router() {
	handler := InitController(Conn())
	router := treemux.New()

	v1 := router.NewGroup("/api/v1")
	v1.GET("/show/:id", handler.Show)
	v1.POST("/store", handler.Store)
	v1.PUT("/update/:id", handler.Update)
	v1.DELETE("/delete/:id", handler.Delete)
}
