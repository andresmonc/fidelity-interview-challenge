package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
)

type Bar struct {
	Uuid string `json:"uuid"`
	Bar  string `json:"bar"`
}

type Bars struct {
	Bars []Bar `json:"bars"`
}

var barsData = Bars{
	Bars: []Bar{
		Bar{Uuid: generateUUID(), Bar: "13"},
		Bar{Uuid: generateUUID(), Bar: "2"},
	},
}

func main() {
	port := "8080"
	fmt.Printf("Starting server at port %s", port)
	http.HandleFunc("/foo/", fooBar)
	startServer(port)
}

func startServer(port string) {
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
	}
}

func fooBar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path,"/foo/")
	switch r.Method {
	case "GET":
		var response []byte
		if(id != ""){
			response = marshalBar(getBarByID(id))
		} else {
			response = marshalBars(barsData)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		isErrPanic(err)
		uuid := addNewBar(unMarshalBar(b))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": ` + uuid + `}`))

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func marshalBars(struc Bars) (btyes []byte) {
	b, err := json.Marshal(struc)
	isErrPanic(err)
	return b
}

//TODO combine with above?
func marshalBar(struc Bar) (btyes []byte) {
	b, err := json.Marshal(struc)
	isErrPanic(err)
	return b
}

func unMarshalBar(bytes []byte) (bar Bar) {
	err := json.Unmarshal(bytes, &bar)
	if err != nil {
		fmt.Println(err)
	}
	return bar
}

func addNewBar(bar Bar) string {
	bar.Uuid = generateUUID()
	barsData.Bars = append(barsData.Bars, bar)
	return bar.Uuid
}

func isErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func generateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	isErrPanic(err)
	return fmt.Sprintf("%x", b)

}

func getBarByID(id string) (bar Bar) {
	barsArray := barsData.Bars
	for i := 0; i < len(barsData.Bars); i++ {
		currentBar := barsArray[i]
		if currentBar.Uuid == id {
			return currentBar
		}
	}
	return
}
