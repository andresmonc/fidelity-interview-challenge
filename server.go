package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Bar represents a bar
type Bar struct {
	// UUID represents unique ID for each bar
	UUID string `json:"UUID"`
	// Bar represents the amount of bars
	Bar int `json:"bar"`
}

// Bars represents a collection of bars
type Bars struct {
	// Bars represents an array of bars
	Bars []Bar `json:"bars"`
}

var barsData = Bars{
	Bars: []Bar{
		Bar{UUID: generateUUID(), Bar: 12},
		Bar{UUID: generateUUID(), Bar: 14},
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
	routeParam := strings.TrimPrefix(r.URL.Path, "/foo/")
	switch r.Method {
	case "GET":
		var response []byte
		if routeParam == "sum" {
			response = []byte(`{"sum":` + fmt.Sprint(sumBar()) + `}`)
		} else if routeParam != "" {
			var bar Bar = getBarByID(routeParam)
			if bar.UUID == "" {
				response = ([]byte(`{"message": "not found"}`))
			} else {
				response = marshalBar(getBarByID(routeParam))
			}
		} else {
			response = marshalBars(barsData)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		isErrPanic(err)
		UUID := addNewBar(unMarshalBar(b))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "` + UUID + `"}`))
	case "DELETE":
		msg := deleteBarByID(routeParam)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "` + msg + `"}`))
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
	bar.UUID = generateUUID()
	barsData.Bars = append(barsData.Bars, bar)
	return bar.UUID
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

func sumBar() int {
	sum := 0
	for i := 0; i < len(barsData.Bars); i++ {
		sum += barsData.Bars[i].Bar
	}
	return sum
}

func deleteBarByID(id string) string {
	indexForID := -1
	for i := 0; i < len(barsData.Bars); i++ {
		currentBar := barsData.Bars[i]
		if currentBar.UUID == id {
			indexForID = i
		}
	}
	if indexForID != -1 {
		barsData.Bars[indexForID] = barsData.Bars[len(barsData.Bars)-1] //copy last element to index
		barsData.Bars = barsData.Bars[:len(barsData.Bars)-1]            // remove last element
		return "Success"
	}
	return "not found"

}

func getBarByID(id string) (bar Bar) {
	for i := 0; i < len(barsData.Bars); i++ {
		currentBar := barsData.Bars[i]
		if currentBar.UUID == id {
			return currentBar
		}
	}
	return
}
