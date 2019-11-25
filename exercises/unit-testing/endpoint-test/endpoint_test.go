package endpoint_test

import (
	"encoding/json"
	"net/http/httptest"
	"net/http"
	"testing"

	"github.com/goinaction/code/chapter9/listing17/handlers"
)

const checkMark = "u\2713"
const ballotX = "\u2717"

func init() {
	handlers.Routes()
}

func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("Should be able to create a request.",
				ballotX, err)
		}
		t.Log("Should be able to create a request.", checkMark)

		rw := httptest.NewRecorder()
	}
}