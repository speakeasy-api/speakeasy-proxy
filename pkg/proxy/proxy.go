package proxy

import (
	"fmt"
	"net/http"

	"github.com/speakeasy-api/speakeasy-go-sdk"
	"github.com/speakeasy-api/speakeasy-proxy/internal/handler"
)

type ProxyConfig struct {
	DownstreamBaseURL string
	Port              string
	APIKey            string
	ApiID             string
	VersionID         string
	OpenAPIDocument   []byte
}

func StartProxy(config ProxyConfig) error {
	handler := handler.NewHandler(config.DownstreamBaseURL)

	sdk := speakeasy.New(speakeasy.Config{
		APIKey:          config.APIKey,
		ApiID:           config.ApiID,
		VersionID:       config.VersionID,
		OpenAPIDocument: config.OpenAPIDocument,
	})

	fmt.Printf("proxy listening on port %s and redirecting to %s\n", config.Port, config.DownstreamBaseURL)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), sdk.Middleware(handler)); err != nil {
		return err
	}

	return nil
}
