package controller

type RunC struct
{
	Code string `json:"code"`
	Stdin string `json:"stdin"`
}

type RunCResponse struct
{
	JobId string `json:"job_id"`
}

func RunCController(req RunC) RunCResponse {
	return RunCResponse{
		JobId: req.Code + req.Stdin,
	}
}

type JobFromId struct
{
	JobId string `json:"job_id"`
}

type JobFromIdResponse struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

func JobFromIDController(req JobFromId) JobFromIdResponse {
	return JobFromIdResponse{
		Stdout: req.JobId,
		Stderr: req.JobId,
	}
}
