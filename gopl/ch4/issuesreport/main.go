package main

import (
	"github.com/eanson023/golearning/gopl/ch4/github"
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range.Items}}-------------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreateAt | daysAgo}} days
{{end}}`

// {{range.Items}} {{end}} 表示一个循环
// | 管道符 前一个操作的结果当做下一个操作的输入 printf在所有模板中是fmt.Sprintf的同义词
// 对于Age来说，第二个操作时daysAge，这个函数使用time.Since将CreateAt转换为已经过去的时间

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// template.Must提供了一个便捷的错误处理方式(panic)
// holy shit!!! 语法看似很诡异 Must方法明明是两个参数 其实调用Parse方法也是返回两个参数，这里刚好需要的类型一模一样！
var report = template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

// // go run main.go repo:golang/go is:open json decoder
func main() {
	// report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err = report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
