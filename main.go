package main

import (
	"github.com/airbrake/gobrake/v5"
	"log"
	"net/http"
)

var airbrake = gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
	ProjectId:   311603,
	ProjectKey:  "4f663f741aa2a901e6619f2e8a9d93b1",
	Environment: "production",
})

func main() {
	Router()
	defer airbrake.Close()
	defer airbrake.NotifyOnPanic()
	log.Println("listen on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
