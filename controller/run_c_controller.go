package controller

type RunC struct
{
	Code string `json:"code"`
	Stdin string `json:"stdin"`
}

type RunCResponse struct
{
	JobId string `json:"jobId"`
}

func RunCController(req RunC) RunCResponse {
	return RunCResponse{
		JobId: "uuid4",
	}
}

type JobFromId struct
{
	JobId string `json:"jobid"`
}

type JobFromIdResponse struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

func JobFromIDController(req JobFromId) JobFromIdResponse {
	return JobFromIdResponse{
		Stdout: "Good",
		Stderr: "Bad",
	}
}
