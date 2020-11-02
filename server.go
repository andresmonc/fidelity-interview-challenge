package main

import (
	"fmt"
	"net/http"
)


type Bar struct {
	uuid   int `json:"uuid"`
	bar    int `json:"bar"`
}
// acting as database
var Bars = []Bar{
	Bar{uuid: 123, bar: 13},
	Bar{uuid: 123, bar: 2},
}


func main() {
	port := "8080"
	fmt.Printf("Starting server at port %s", port)
	startServer(port)
}

func startServer(port string) {
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
	}
}

