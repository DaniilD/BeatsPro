package controllers

import (
	"BeatsPro/internal/requests"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type TagController struct {
}

func NewTagController() *TagController {
	return &TagController{}
}

func (tagController *TagController) CreateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var CreateTagRequest requests.CreateTagRequest

	if err := json.NewDecoder(r.Body).Decode(&CreateTagRequest); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	stringRequest := string(body)
	fmt.Fprintln(os.Stdout, stringRequest)
}
