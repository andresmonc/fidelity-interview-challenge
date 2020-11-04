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

func Test_addNewBar(t *testing.T) {
	barsLen := len(barsData.Bars)
	bar := Bar{UUID: "223456789", Bar: 1}
	addNewBar(bar)
	assertEquals(barsLen+1, len(barsData.Bars))
}

func Test_getBarByID(t *testing.T) {
	assertEquals(99, getBarByID("123456789").Bar)
}

func Test_marshalBar(t *testing.T) {
	marshalBar(barsData.Bars[0])
}

func Test_marshalBars(t *testing.T) {
	marshalBars(barsData)
}

func Test_unMarshalBars(t *testing.T) {
	bar := Bar{UUID: "223456789", Bar: 1}
	bytes := marshalBar(bar)
	newBar := unMarshalBar([]byte(bytes))
	assertEquals(1, newBar.Bar)
}
