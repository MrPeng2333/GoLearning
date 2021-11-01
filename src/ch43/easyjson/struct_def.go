package jsontest

//easyjson
type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//easyjson
type JobInfo struct {
	Skills []string `json:"skills"`
}

//easyjson
type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}
