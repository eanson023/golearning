package goquery

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
	"testing"
	"unicode/utf8"
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
	dom.Find("#divShow1>table").First().Find("tbody>tr").Each(func(i int, sl *goquery.Selection) {
		if sl.HasClass("rnm") {
			t.Log("ashaioashia")
		}
		t.Log(sl.Html())
		t.Log("***********************")
		// sl.Each(func(j int, s2 *goquery.Selection) {
		// 	t.Log(s2.Html())
		// })
	})
}

func TestTT(t *testing.T) {
	var scriptText string = "alert(hhhhhh,looj)"
	errorMsg := scriptText[strings.Index(scriptText, "(")+1 : strings.LastIndex(scriptText, ")")]
	t.Log(errorMsg)
}

func Test2(t *testing.T) {
	str := "所选学分24.50；获得学分24.50；重修学分。"
	split := strings.Split(str, "；")
	for _, v := range split {
		t.Log(v[4*3:])
		t.Log(utf8.RuneCountInString(v))
	}
}

func Test3(t *testing.T) {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	if err := f.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`); err != nil {
		t.Error(err)
		return
	}
	// Save xlsx file by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		t.Error(err)
	}
}
