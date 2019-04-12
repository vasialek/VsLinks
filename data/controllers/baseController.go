package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vasialek/VsLinks/models"
)

func reportError(w http.ResponseWriter, msg string, err error) {
	fmt.Println(err)
	resp := models.Response{
		Status:  false,
		Message: msg,
	}
	w.WriteHeader(http.StatusNotAcceptable)
	j, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Write(j)
}
