package main

import (
	"fmt"
	"net/http"
)

type Bar struct {
	uuid int `json:"uuid"`
	bar  int `json:"bar"`
}

// acting as database
var Bars = []Bar{
	Bar{uuid: 123, bar: 13},
	Bar{uuid: 123, bar: 2},
}

func main() {
	port := "8080"
	fmt.Printf("Starting server at port %s", port)
	http.HandleFunc("/", home)
	startServer(port)
}

func startServer(port string) {
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
        w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
	}
}


func  getBarById(id int) (bar Bar){
	for i := 0; i < len(Bars); i++ {
		currentBar := Bars[i]
		if currentBar.uuid == id {
			return currentBar
		}
	}
	return
}
