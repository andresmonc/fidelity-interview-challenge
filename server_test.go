package main

import (
	"fmt"
	"testing"
)

var mockBar = Bars{
	Bars: []Bar{
		Bar{UUID: generateUUID(), Bar: 99},
		Bar{UUID: generateUUID(), Bar: 1},
	},
}

func Test_sumBar(t *testing.T) {
	assertEquals(100, sumBar(mockBar))
}

func assertEquals(expected int, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("%d not equal %d", expected, actual))
	}
}
