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
