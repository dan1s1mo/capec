package myroutines

import (
	fe "capec/fileeditor"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type FileMeta struct {
	Count int32  `json:"count"`
	Name  string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type FileWorker struct {
	port      string
	index     int
	filesMeta []FileMeta
}

func (fw *FileWorker) Init(port, files string) error {
	http.HandleFunc("/files-update", fw.handler)
	fmt.Printf("Listening on port %s...", port)
	fw.port = port
	jsonFile, err := os.Open(files)
	if err != nil {
		return err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &fw.filesMeta)
	if err != nil {
		return err
	}
	return nil
}

func (fw *FileWorker) Run() {
	http.ListenAndServe(":"+fw.port, nil)
}

func (fw *FileWorker) handler(w http.ResponseWriter, r *http.Request) {
	fw.index = (fw.index + 1) % len(fw.filesMeta)
	meta := fw.filesMeta[fw.index]
	fileErr := fe.ModifyFiles(meta.Name, int(meta.Count))
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
