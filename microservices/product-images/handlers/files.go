package handlers

import (
	"github.com/eanson023/golearning/microservices/product-images/files"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type Files struct {
	log   *log.Logger
	store files.Storage
}

func NewFiles(s files.Storage, l *log.Logger) *Files {
	return &Files{l, s}
}

// UploadREST 标准的RESTful风格上传文件
func (f *Files) UploadREST(rw http.ResponseWriter, req *http.Request) {
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

// UploadMulipart 解析multipart/form
// <form action="http://localhost:8080" method="POST" enctype="multipart/form-data">
//   <p>fileid:<input type="text" name="id" value=""></p>
//   <p><input type="file" name="file"></p>
//   <p><input type="submit" value="Submit"></p>
// </form>
// 上传文件会有name和filename两个属性
func (f *Files) UploadMulipart(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Println("[ERROR] Bad Request ", err.Error())
		http.Error(rw, "Expected multiform data", http.StatusBadRequest)
		return
	}
	// 表达值
	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		f.log.Println("[ERROR] Bad Request ", err.Error())
		http.Error(rw, "Expected integer id", http.StatusBadRequest)
		return
	}
	f.log.Printf("Process form for id:%d", id)
	muliFile, mh, err := req.FormFile("file")
	if err != nil {
		f.log.Println("[ERROR] Bad Request ", err.Error())
		http.Error(rw, "Expected multiform data", http.StatusBadRequest)
		return
	}
	// 保存文件
	f.saveFile(req.FormValue("id"), mh.Filename, rw, muliFile)

}

func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.log.Println("Invalid path", uri)
	http.Error(rw, "Invalid file path shouldbe in the format /[id]/[filepath]", http.StatusBadRequest)
}

func (f *Files) saveFile(id, filename string, rw http.ResponseWriter, reader io.ReadCloser) {
	filePath := filepath.Join(id, filename)
	err := f.store.Save(filePath, reader)
	if err == files.ExceedMustFileSizeError {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	} else if err != nil {
		f.log.Println("Unable to save file error:", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
