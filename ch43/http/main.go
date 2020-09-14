package main

import (
	"encoding/json"
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

func main() {
	http.HandleFunc("/api/users", func(writer http.ResponseWriter, request *http.Request) {
		bs, _ := json.Marshal(user)
		writer.Header()["Access-Control-Allow-Origin"] = []string{"*"}
		writer.Write(bs)
	})
	http.ListenAndServe(":8080", nil)
}
