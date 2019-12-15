package controllers

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) SetValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	return
}