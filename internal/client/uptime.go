package client

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/atharv3221/uptimecheck/internal/response"
)

func UptimeCheck(request string) {
	resp, err := http.Get(request)

	if err != nil {
		if err, ok := err.(*url.Error); ok {
			output := response.URLError(err)
			slog.Error("error occured", "response", output)
			return
		}
		output := response.URLErr(err, resp)
		slog.Error("error occured", "response", output)
		return
	}
	slog.Info("Successful", "response", response.GenerateResponse(resp))
}
