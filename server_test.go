package main

import (
	"testing"
)

func init() {
	barsData = Bars{
		Bars: []Bar{
			Bar{UUID: "123456789", Bar: 99},
			Bar{UUID: "223456789", Bar: 1},
		},
	}
}

func Test_sumBar(t *testing.T) {
	assertEquals(100, sumBar())
}

func Test_generateUUID(t *testing.T) {
	assertEquals(len(generateUUID()), 32)
}

func Test_deleteBarByID(t *testing.T) {
	assertTrue(deleteBarByID("") == "not found")
	assertTrue(deleteBarByID("223456789") == "Success")
}
