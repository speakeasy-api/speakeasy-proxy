package handler

import (
	"io"
	"net/http"
	"net/url"
)

type Handler struct {
	redirectURL string
}

func NewHandler(redirectURL string) *Handler {
	return &Handler{
		redirectURL: redirectURL,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	downstreamURL, err := url.JoinPath(h.redirectURL, r.URL.Path)
	if err != nil {
		handleError(w, err)
		return
	}

	downstreamRequest, err := http.NewRequest(r.Method, downstreamURL, r.Body)
	if err != nil {
		handleError(w, err)
		return
	}
	downstreamRequest.Header = r.Header

	c := http.Client{}
	resp, err := c.Do(downstreamRequest)
	if err != nil {
		handleError(w, err)
		return
	}

	for name, headers := range resp.Header {
		for _, header := range headers {
			w.Header().Add(name, header)
		}
	}

	w.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(w, resp.Body); err != nil {
		handleError(w, err)
		return
	}
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := io.WriteString(w, err.Error()); err != nil {
		panic(err)
	}
}
