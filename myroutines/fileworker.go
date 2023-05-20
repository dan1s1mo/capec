package myroutines

import (
	fe "capec/fileeditor"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type FileWorker struct {
	port string
}

func (fw *FileWorker) Init(port string) {
	http.HandleFunc("/files-update", handler)
	fmt.Printf("Listening on port %s...", port)
	fw.port = port
}

func (fw *FileWorker) Run() {
	http.ListenAndServe(":"+fw.port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fileErr := fe.ModifyFiles(".\\test", 1)
	if fileErr != nil {
		http.Error(w, fileErr.Error(), http.StatusInternalServerError)
		return
	}
	res := Response{
		Message: "Done",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
