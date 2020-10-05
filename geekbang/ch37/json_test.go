package ch37

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

// 内置json解析
func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	// 反序列化
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(e.BasicInfo, e.JobInfo)
	// 序列化
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}

// easyjson代码生成的方式测试
func TestEasyJson(t *testing.T) {
	e := new(Employee)
	e.UnmarshalJSON([]byte(jsonStr))
	t.Log(e.BasicInfo, e.JobInfo)
	if value, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		t.Log(string(value))
	}
}

// 性能测试
func BenchmarkEmbeddedJson(b *testing.B) {
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		// 反序列化
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		// 序列化
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyJson(b *testing.B) {
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		e.UnmarshalJSON([]byte(jsonStr))
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}

/*
easyjson大概快了4倍
BenchmarkEmbeddedJson
BenchmarkEmbeddedJson-8   	  365155	      3183 ns/op	     568 B/op	      10 allocs/op
BenchmarkEasyJson
BenchmarkEasyJson-8       	 1259168	       933 ns/op	     252 B/op	       5 allocs/op
*/
