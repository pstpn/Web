package http

type StatusBadRequest struct {
	Error string `json:"error" example:"Incorrect request body date of birth"`
}

type StatusInternalServerError struct {
	Error string `json:"error" example:"Failed to register employee"`
}

type StatusUnauthorized struct {
	Error string `json:"error" example:"Cant login employee"`
}

type StatusNotFound struct {
	Error string `json:"error" example:"Failed to get info card document"`
}
