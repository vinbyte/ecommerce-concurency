package domain

// SuccessBaseResponse used as base response form if the code is 2xx
type SuccessBaseResponse struct {
	Data interface{} `json:"data"`
}

// ErrorDataResponse is item of ErrorBaseResponse.Errors array
type ErrorDataResponse struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

// ErrorBaseResponse used as base response form if the code is 4xx and 5xx
type ErrorBaseResponse struct {
	Error struct {
		Code    int                 `json:"code"`
		Message string              `json:"message"`
		Errors  []ErrorDataResponse `json:"errors"`
	} `json:"error"`
}
