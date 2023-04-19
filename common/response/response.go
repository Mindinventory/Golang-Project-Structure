package response

type NoContent struct {
}

type StatusResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Cause string `json:"cause"`
}
