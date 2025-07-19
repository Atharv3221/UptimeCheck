package structure

import "time"

// WebsiteStatus represents whether the site is up or down.
type WebsiteStatus string

const (
	Up   WebsiteStatus = "up"
	Down WebsiteStatus = "down"
)

// WebsiteResponse represents the result of a website uptime check.
type WebsiteResponse struct {
	URL            string        `json:"url"`
	Status         WebsiteStatus `json:"status"`
	HTTPStatus     int           `json:"http_status"`
	ResponseTimeMs int           `json:"response_time_ms"`
	CheckedAt      time.Time     `json:"checked_at"`
}

// Methods for interface and usage convenience
func (w WebsiteResponse) GetStatus() int {
	return w.HTTPStatus
}

func (w WebsiteResponse) GetURL() string {
	return w.URL
}
