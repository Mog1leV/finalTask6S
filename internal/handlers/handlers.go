package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error while receiving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	dataFile, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	convert := service.ConverterTo(string(dataFile))
	newFileName := time.Now().UTC().String() + ".txt"

	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "error creating file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = io.WriteString(newFile, convert)
	if err != nil {
		http.Error(w, "error writing to file", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(convert))
}
