// gif动画
package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/eanson023/golearning/gopl/ch1/lissajous/lissa"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissa.Lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":8080", nil))
		return
	}
	lissa.Lissajous(os.Stdout)
}
