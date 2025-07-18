package response

import (
	"net/http"
	"net/url"
)

type status string

const (
	Successful status = "Request was successful"
	Failed     status = "Request was failed"
)

type Response struct {
	statusCode     int
	res            *http.Response
	responseStatus status
}

func URLError(err *url.Error) Response {
	response := Response{
		res:            nil,
		responseStatus: Failed,
	}
	return response
}

func URLErr(err error, resp *http.Response) Response {
	response := Response{
		statusCode:     resp.StatusCode,
		res:            resp,
		responseStatus: Failed,
	}
	return response
}

func GenerateResponse(resp *http.Response) Response {
	response := Response{
		statusCode:     resp.StatusCode,
		res:            resp,
		responseStatus: Successful,
	}
	return response
}
