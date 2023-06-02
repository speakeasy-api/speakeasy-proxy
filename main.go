package main

import (
	"os"

	"github.com/speakeasy-api/speakeasy-proxy/pkg/proxy"
)

const downstreamBaseURL = "http://localhost:8001"

func main() {
	doc, err := os.ReadFile("./httpbin.yaml")
	if err != nil {
		panic(err)
	}

	if err := proxy.StartProxy(proxy.ProxyConfig{
		DownstreamBaseURL: downstreamBaseURL,
		Port:              "3333",
		APIKey:            "eyJpZCI6ImNsaWN6ejQxeDAwMDAzNTZkZG1zZGx1aW4iLCJuYW1lIjoicHJveHktdGVzdCIsIndvcmtzcGFjZV9pZCI6ImNsNzY0cG5peDAwMDAzNTZkNWl3eTB5MTEiLCJjcmVhdGVkX2F0IjoiMjAyMy0wNi0wMVQxMDozMjoyNi45NDlaIiwiYWxnIjoiU0hBUkVEX1NFQ1JFVCIsInNlY3JldCI6IkpVTXlKVGt3SlRBMUpVTXpKVUUyUnlWRE15VTVOQ1ZETlNVNU1pVTJNQ1V5UXlWRE15VkNNMU1uSlVNeUpVSkRKVU16SlVJMUpVTXpKVGswSlVNekpVSkdjaVZETWlWQ01pVkZNaVU0TUNVNVFTVkRNaVZCUXlWRFFpVTROaVZETWlVNU1DVkRNaVZDUVNWRE15VTVOQ1ZGTWlVNE1DVkNNR0lsUXpNbFFqa2xReklsUWprbFF6SWxRa1FsTVRVbFF6SWxRVVVsUlRJbE9ESWxRVU1sUXpNbFFrUWxNVEVsTWpZbFF6TWxPVVZCSlVNeUpVSTVKVEl3VWxnbFF6TWxPRU1sTURnbE1EUmxKVEJFSlVVeUpUZzBKVUV5SlRGQkpVTXpKVUpHSlRFMEpUZENKVU16SlRoQ0pVTXpKVUkxSlRBM0pVTTFKVUV3SlVNeUpVSTFKVFF3TWlVeU1pVXdOeVZETWlWQ015VkRNaVU0TVNWRE15VkNSU1ZETXlWQlFTVXdNQT09IiwiY3JlYXRlZF9ieSI6IjVlNDI5ZWM4LTcwYWMtNDlmYi1iMTc0LTJlMjhhYzlmMzNmNSJ9",
		ApiID:             "proxy-test",
		VersionID:         "0.0.1",
		OpenAPIDocument:   doc,
	}); err != nil {
		panic(err)
	}
}
