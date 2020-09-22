package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

// User user
type User struct {
	Username string   `json:"username"`
	Age      int      `json:"age"`
	Hobby    []string `json:"hobby"`
}

var user *User = new(User)

func init() {
	user.Username = "eanson"
	user.Age = 18
	user.Hobby = []string{"敲代码", "打篮球", "睡觉"}
}
func GetUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	byes, _ := json.Marshal(user)
	writer.Header()["Access-Control-Allow-Origin"] = []string{"*"}
	writer.Write(byes)
}

func AddUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	body := request.Body
	bs, _ := ioutil.ReadAll(body)
	user2 := new(User)
	json.Unmarshal(bs, user2)
	fmt.Println(user2)
	writer.Header()["Content-Type"] = []string{"application/json"}
	writer.Write([]byte("{\"msg\",\"ok\"}"))
}

func main() {
	router := httprouter.New()
	router.GET("/api/user", GetUsers)
	router.POST("/api/user", AddUser)
	http.ListenAndServe(":8080", router)
}
