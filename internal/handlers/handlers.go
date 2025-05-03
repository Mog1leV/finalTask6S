package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("myFile")
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
	extension := filepath.Ext(header.Filename)
	newFileName := time.Now().UTC().Format("2006-01-02_15-04-05") + extension

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
