package structure

import "time"

// HTTPMethod represents HTTP verbs.
type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
)

// APIStatus represents success or failure of the API response.
type APIStatus string

const (
	Failed     APIStatus = "failed"
	Successful APIStatus = "successful"
)

// APIResponse is the main struct representing an API test/check result.
type APIResponse struct {
	URL            string            `json:"url"`
	Method         HTTPMethod        `json:"method"`
	Status         APIStatus         `json:"status"`
	HTTPStatus     int               `json:"http_status"`
	ResponseTimeMs int               `json:"response_time_ms"`
	Headers        map[string]string `json:"headers"`
	Body           string            `json:"body"`
	CheckedAt      time.Time         `json:"checked_at"`
}

// Methods to satisfy some interface (if needed).
func (a APIResponse) GetStatus() int {
	return a.HTTPStatus
}

func (a APIResponse) GetURL() string {
	return a.URL
}

func (a APIResponse) GetBody() string {
	return a.Body
}
