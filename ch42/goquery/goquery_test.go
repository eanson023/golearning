package goquery

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestGoquery(t *testing.T) {
	file, err := os.Open("../../ch40/server/index.htm")
	if err != nil {
		t.Fatal(err)
	}
	dom, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal(err)
	}
	str := dom.Find("script").Last().Text()
	assert.True(t, strings.Contains(str, "验证码不正确"))
}
