package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var width = flag.Int("w", 256, "hash width(256 or 512)")

func main() {
	flag.Parse()
	// 闭包牛批
	var funcation func(b []byte) []byte
	switch *width {
	case 256:
		funcation = func(b []byte) []byte {
			sha := sha256.Sum256(b)
			return sha[:]
		}
	case 512:
		funcation = func(b []byte) []byte {
			sha := sha512.Sum512(b)
			return sha[:]
		}
	default:
		log.Fatal("Unexpected width specified.")
	}
	fmt.Fprintln(os.Stdout, "please input data")
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "%x\n", funcation(b))
}
