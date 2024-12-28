package util

type Response struct {
    Message string `json:"message"`
}

func ErrorResponse(message string) Response {
    return Response{Message: message}
}
