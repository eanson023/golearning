package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	var letterCount, numberCount, unicodeCount int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letterCount++
		} else if unicode.IsNumber(r) {
			numberCount++
		} else {
			unicodeCount++
		}
	}
	fmt.Print("rune\tcount\n")
	fmt.Printf("letters:%d\tnumbers:%d\tunicodes:%d\n", letterCount, numberCount, unicodeCount)
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
