package gateway

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	httpSvcAddress = "http://http-service:8080"
)

func HttpCreate(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post(httpSvcAddress+"/quests", "application/json", r.Body)
	if err != nil {
		handleUnreachable(err, w)
		return
	}

	defer resp.Body.Close()
	message, _ := ioutil.ReadAll(resp.Body)

	_, err = w.Write(message)
}

func HttpGetAll(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(httpSvcAddress + "/quests")
	if err != nil {
		handleUnreachable(err, w)
		return
	}

	defer resp.Body.Close()
	jsonBody, _ := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBody)
}

func HttpDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", httpSvcAddress+"/quests/"+vars["id"], nil)
	if err != nil {
		log.Print("HTTP handler: Creating DELETE request failed.")
	}

	resp, err := client.Do(req)
	if err != nil {
		handleUnreachable(err, w)
		return
	}

	defer resp.Body.Close()
	message, _ := ioutil.ReadAll(resp.Body)

	_, err = w.Write(message)
}

func handleUnreachable(err error, w http.ResponseWriter) {
	log.Printf("Unreachable: %v", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
