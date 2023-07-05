package proxy

import (
	"fmt"
	"net/http"

	"github.com/speakeasy-api/speakeasy-go-sdk"
	"github.com/speakeasy-api/speakeasy-proxy/internal/handler"
)

type Options struct {
	DownstreamBaseURL string
	Port              string
	APIKey            string
	ApiID             string
	VersionID         string
	OpenAPIDocument   []byte
}

func StartProxy(opts Options) error {
	handler := handler.NewHandler(opts.DownstreamBaseURL)

	sdk := speakeasy.New(speakeasy.Config{
		APIKey:          opts.APIKey,
		ApiID:           opts.ApiID,
		VersionID:       opts.VersionID,
		OpenAPIDocument: opts.OpenAPIDocument,
	})

	fmt.Printf("proxy listening on port %s and redirecting to %s\n", opts.Port, opts.DownstreamBaseURL)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", opts.Port), sdk.Middleware(handler)); err != nil {
		return err
	}

	return nil
}
