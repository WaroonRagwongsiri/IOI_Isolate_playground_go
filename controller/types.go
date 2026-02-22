package controller

type RunC struct {
	Code  string `json:"code"`
	Stdin string `json:"stdin"`
}

type RunCResponse struct {
	JobId string `json:"job_id"`
}

type JobFromId struct {
	JobId string `json:"job_id"`
}

type JobFromIdResponse struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}
