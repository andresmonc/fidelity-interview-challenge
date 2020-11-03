package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Bar struct {
	Uuid string `json:"uuid"`
	Bar  string `json:"bar"`
}

// acting as database
var Bars = []Bar{
	Bar{Uuid: "123", Bar: "13"},
	Bar{Uuid: "1233", Bar: "2"},
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
		response := marshalBar(Bars)
        w.WriteHeader(http.StatusOK)
		w.Write(response)
	default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
	}
}

func marshalBars(struc []Bar) (btyes []byte) {
	b, err := json.Marshal(struc)
	fmt.Println(struc)
	if err != nil {
		panic(err)
	}
	return b
}

//TODO combine with above?
func marshalBar(struc Bar) (btyes []byte) {
	b, err := json.Marshal(struc)
	fmt.Println(struc)
	if err != nil {
		panic(err)
	}
	return b
}

func  getBarByID(id string) (bar Bar){
	for i := 0; i < len(Bars); i++ {
		currentBar := Bars[i]
		if currentBar.Uuid == id {
			return currentBar
		}
	}
	return
}
