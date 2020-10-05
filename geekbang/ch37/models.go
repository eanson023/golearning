package ch37

// 通过fieldTag来映射json关系
// 反射影响性能
type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo *BasicInfo `json:"basic_info"`
	JobInfo   *JobInfo   `json:"job_info"`
}
