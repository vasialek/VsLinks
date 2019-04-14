package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vasialek/VsLinks/models"
)

func sendDataResponse(w http.ResponseWriter, data interface{}) error {
	ba, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ba)
	return nil
}

func reportError(w http.ResponseWriter, msg string, err error) {
	if err != nil {
		fmt.Println(err)
	}
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
