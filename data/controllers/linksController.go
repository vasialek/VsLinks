package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vasialek/VsLinks/data"
	"github.com/vasialek/VsLinks/models"
)

// LinksController to be class for lins controller
type LinksController struct{}

// NewLinksController returns instance of LinksController
func NewLinksController() *LinksController {
	return &LinksController{}
}

// GetLinks returns list of links for current user
func (lc *LinksController) GetLinks(w http.ResponseWriter, rq *http.Request) {
	list, err := data.GetAllLinks()
	if err != nil {
		reportError(w, "Error getting Links from database.", err)
		return
	}

	b, err := json.Marshal(&list)
	if err != nil {
		reportError(w, "Error JSONing", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// CreateLink creates link in DB from POSTed JSON
func (lc *LinksController) CreateLink(w http.ResponseWriter, rq *http.Request) {
	log.Println("Going to create link...")
	var model models.Link
	err := json.NewDecoder(rq.Body).Decode(&model)
	if err != nil {
		log.Printf("createLink: %s\n", err)
		reportError(w, "Error deserializing Link to be created.", err)
		return
	}

	log.Printf("  link to create: %s\n", model.URL)
	err = data.CreateLink(model)
	if err != nil {
		reportError(w, "Can't save Link in database.", err)
		return
	}

	resp := models.Response{
		Message: fmt.Sprintf("New Link `%s` was created", model.Title),
		Status:  true,
	}
	ba, err := json.Marshal(&resp)
	if err != nil {
		reportError(w, "Can't serialize positive JSON response", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(ba)
}
