package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
}

type httpChecker struct {
}

func (h httpChecker) Check(ctx context.Context, name string) (status bool, err error) {
	response, err := http.Get(name)
	if err != nil || response.StatusCode != 200 {
		return false, err
	}
	return true, nil

}

type websites struct {
	Data map[string]bool `json:"data"`
}

var res = &websites{
	Data: make(map[string]bool),
}

var websiteStatus = map[string]string{}

func main() {
	var mutex sync.Mutex
	go func() {
		for {
			mutex.Lock()
			defer mutex.Unlock()
			updateStatus()
			time.Sleep(1 * time.Minute)
		}
	}()

	mux := http.DefaultServeMux

	mux.HandleFunc("POST /websites", postWebsitesList)
	mux.HandleFunc("GET /websites", handleStatus)

	fmt.Print("Server running on port 3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		fmt.Println(err)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	val := params.Get("name")

	if val != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		var a httpChecker
		status, err := a.Check(ctx, val)

		if err != nil {
			websiteStatus[val] = "DOWN"
			http.Error(w, "Error getting response", http.StatusBadRequest)
			return
		}
		if !status {
			websiteStatus[val] = "DOWN"
		} else {
			websiteStatus[val] = "UP"
		}

		handleResponse(w, map[string]string{val: websiteStatus[val]}, http.StatusOK)
	} else {
		updateStatus()
		handleResponse(w, websiteStatus, http.StatusOK)
	}
}

func postWebsitesList(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	websitesList := &struct {
		Data []string `json:"data"`
	}{Data: nil}

	err = json.Unmarshal(body, &websitesList)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	for _, val := range websitesList.Data {
		if _, ok := res.Data[val]; !ok {
			res.Data[val] = false
		}
	}

	handleResponse(w, res.Data, http.StatusOK)
}

func updateStatus() {
	for url := range res.Data {

		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
		defer cancel()
		var a httpChecker
		status, err := a.Check(ctx, url)
		if err != nil {
			res.Data[url] = false
			websiteStatus[url] = "DOWN"
			fmt.Println("Error getting response from website", err)
			continue
		}
		if status {
			res.Data[url] = true
			websiteStatus[url] = "UP"

		} else {
			res.Data[url] = false
			websiteStatus[url] = "DOWN"
		}
	}

}

func handleResponse(w http.ResponseWriter, message any, statusCode int) {

	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Error converting response to json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
