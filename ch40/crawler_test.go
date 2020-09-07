package ch40

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

const URL string = "http://127.0.0.1:8080"

func TestSimpleReq(t *testing.T) {
	if resp, err := http.Get(URL); err != nil {
		t.Fatal(err)
	} else {
		reader := resp.Body
		printResponse(reader)
	}
}
func TestPost(t *testing.T) {
	file, erro := os.Open("./photo.jpg")
	if erro != nil {
		t.Fatal(erro)
	}
	// reader := bufio.NewReader(file)
	if resp, err := http.Post(URL+"/upload", "image/jpeg", file); err != nil {
		t.Fatal(err)
	} else {
		printResponse(resp.Body)
	}

}
func printResponse(reader io.Reader) {
	bufReader := bufio.NewReader(reader)
	for {
		str, erro := bufReader.ReadString('\n')
		if erro != nil && erro != io.EOF {
			panic(erro)
		}
		fmt.Printf(str)
		if erro == io.EOF {
			fmt.Printf("\nread end\n")
			break
		}
	}
}
