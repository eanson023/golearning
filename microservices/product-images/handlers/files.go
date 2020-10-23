package handlers

import (
	"github.com/eanson023/golearning/microservices/product-images/files"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

type Files struct {
	log   *log.Logger
	store files.Storage
}

func NewFiles(s files.Storage, l *log.Logger) *Files {
	return &Files{l, s}
}

func (f *Files) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fn := vars["filename"]
	f.log.Printf("Handle POST id:%s\tfilename:%s\n", id, fn)

	// check that the filepath is a valid name and file
	if id == "" || fn == "" {
		f.invalidURI(req.URL.String(), rw)
		return
	}
	f.saveFile(id, fn, rw, req.Body)
}

func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.log.Println("Invalid path", uri)
	http.Error(rw, "Invalid file path shouldbe in the format /[id]/[filepath]", http.StatusBadRequest)
}

func (f *Files) saveFile(id, filename string, rw http.ResponseWriter, body io.ReadCloser) {
	filePath := filepath.Join(id, filename)
	err := f.store.Save(filePath, body)
	if err == files.ExceedMustFileSizeError {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	} else if err != nil {
		f.log.Println("Unable to save file error:", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
