package rest_err

type RestErr struct {
	Message string `json:"message"`
	Err string `json:"error`
	Code int `json:"code`
	Causes []Causes `json:"causes"`
}

type Causes struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr{
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
		Causes: causes,
	}
}

func NewInteralServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad_request",
		Code: http.StatusInternalServerError,
	}
}