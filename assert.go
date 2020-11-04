package main

import "fmt"

func assertEquals(expected int, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("%d not equal %d", expected, actual))
	}
}

func assertTrue(boolean bool) {
	if !boolean {
		panic("received boolean is false")
	}
}
