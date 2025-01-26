package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type websites struct {
	Data map[string]bool `json:"data"`
}

var res = &websites{
	Data: make(map[string]bool),
}

var websiteStatus = map[string]string{}

func main() {
	go func() {
		for {
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
		getWebsiteStatus(val, w)
	} else {
		getStatus(w)
	}
}

func getWebsiteStatus(val string, w http.ResponseWriter) {

	response, err := http.Get(val)

	if err != nil {
		http.Error(w, "Error getting response", http.StatusBadRequest)
		return
	}
	if response.StatusCode != 200 {
		websiteStatus[val] = "DOWN"
		return
	}

	websiteStatus[val] = "UP"
	handleResponse(w, map[string]string{val: websiteStatus[val]}, http.StatusOK)

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
func getStatus(w http.ResponseWriter) {

	updateStatus()
	handleResponse(w, websiteStatus, http.StatusOK)

}
func updateStatus() {
	for url := range res.Data {
		
		response, err := http.Get(url)
		if err != nil {
			res.Data[url] = false
			websiteStatus[url] = "DOWN"
			fmt.Println("Error getting response from website", err)
			return
		}
		if response.StatusCode == 200 {
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
