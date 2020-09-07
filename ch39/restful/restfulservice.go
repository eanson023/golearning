package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var employeeDB map[string]*Employee

func init() {
	employeeDB = make(map[string]*Employee, 3)
	employeeDB["Mike"] = &Employee{"e-1", "Mike", 35}
	employeeDB["Eanson"] = &Employee{"e-2", "Mike", 18}
	employeeDB["Rose"] = &Employee{"e-3", "Rose", 20}
}

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello World"))
}

func GetEmployeeByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	qName := params.ByName("name")
	var (
		ok       bool
		info     *Employee
		infoJson []byte
		err      error
	)
	if info, ok = employeeDB[qName]; !ok {
		hearder := writer.Header()
		hearder["Location"] = []string{"https://www.baidu.com"}
		// fmt.Fprintf(writer, "{\"error\":\"404 Not Found\"}")
		writer.WriteHeader(302)
		return
	}
	if infoJson, err = json.Marshal(info); err != nil {
		fmt.Fprintf(writer, "{\"error\":\"%s\"}", err)
		return
	}
	writer.Write(infoJson)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/employee/:name", GetEmployeeByName)
	http.ListenAndServe(":8080", router)
}
