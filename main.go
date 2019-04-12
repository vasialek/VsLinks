package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aleksej.vasinov/visma-links/data"
	"github.com/aleksej.vasinov/visma-links/models"
	"github.com/gorilla/mux"
)

// var links = make(map[string]models.Link)

func main() {
	fmt.Printf("Start working on %s...\n", models.Settings.GetEnvironment())
	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = ":8079"
	} else {
		port = ":" + port
	}

	// uid, _ := uuid.NewV4()
	// links[uid.String()] = models.Link{Title: "New", Url: "http://www.golang.com"}

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/links", getLinks).Methods("GET")
	r.HandleFunc("/links", createLink).Methods("POST")

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	fmt.Println("Going to listen on", port)
	server.ListenAndServe()
}
func createLink(w http.ResponseWriter, rq *http.Request) {
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

func getLinks(w http.ResponseWriter, request *http.Request) {
	// var list []models.Link

	// for _, value := range links {
	// 	list = append(list, value)
	// }

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

func indexHandler(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("Welcome to MemoUs API server"))
}

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
