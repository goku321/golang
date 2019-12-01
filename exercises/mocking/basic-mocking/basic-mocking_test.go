package main

import (
	"golang/exercises/mocking/basic-mocking/mocks"
	"testing"
)

func TestGetString(t *testing.T) {
	expected := "test string"

	mockStringer := &mocks.Stringer{}
	mockStringer.On("String").Return("test string")
	y := mockStringer.String()
	if y != expected {
		t.Errorf("%s != %s", y, expected)
	}
}

func TestDisplay(t *testing.T) {
	m := &mocks.Stringer{}
	m.On("PrintString", "hi there").Return(nil)
	Display(m)
	Display(m)
	Display(m)
	m.AssertNumberOfCalls(t, "PrintString", 3)
}
