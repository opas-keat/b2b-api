package status_code

import "fmt"

// StatusCode
const (
	Success    = 200
	BadRequest = 400
	NotFound   = 401
	Duplicate  = 402
	Forbidden  = 403
	Internal   = 500
)

var statusCodeMessage = map[int]string{
	200: "Success",
	400: "Bad Request",
	401: "Not Found",
	402: "Duplicate",
	403: "Forbidden",
	500: "Internal",
}

// StatusCodeMessage return message
func StatusCodeMessage(code int) string {
	if msg, ok := statusCodeMessage[code]; ok {
		return msg
	}
	return fmt.Sprintf("undefined message for status code: %d", code)
}
