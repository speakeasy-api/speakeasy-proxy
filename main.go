package main

import (
	"os"

	"github.com/speakeasy-api/speakeasy-proxy/internal/config"
	"github.com/speakeasy-api/speakeasy-proxy/pkg/proxy"
	"github.com/speakeasy-api/speakeasy/pkg/merge"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	docPath := cfg.OpenAPIDocs[0]

	if len(cfg.OpenAPIDocs) > 1 {
		docPath = "./proxy-merged.yaml"
		if err := merge.MergeOpenAPIDocuments(cfg.OpenAPIDocs, docPath); err != nil {
			panic(err)
		}
	}

	doc, err := os.ReadFile(docPath)
	if err != nil {
		panic(err)
	}

	if err := proxy.StartProxy(proxy.Options{
		DownstreamBaseURL: cfg.DownstreamBaseURL,
		Port:              cfg.Port,
		APIKey:            cfg.APIKey,
		ApiID:             cfg.ApiID,
		VersionID:         cfg.VersionID,
		OpenAPIDocument:   doc,
	}); err != nil {
		panic(err)
	}
}
