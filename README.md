# speakeasy-proxy

## Overview

The Speakeasy Proxy is a simple reverse proxy that allows capturing of API traffic for use with the Speakeasy Platform.

The proxy is currently designed to be single use, and will only capture HTTP traffic for a single API, this means that it should not be used as a replacement for traditional reverse proxies or API gateways and instead used in conjunction with them.

The proxy requires an OpenAPI document to be provided which helps to associate the captured traffic with the API.

## Usage

### Configuration

The proxy is configured using a YAML file. The following is an example configuration:

```yaml
downstreamBaseURL: http://localhost:8001
apiID: proxy-test
versionID: "0.0.1"
openAPIDocs:
  - ./httpbin.yaml
```

The available configuration options are:

| Option              | Description                                                                                                                                                                                                                          | Required           | Default              | Location                                      |
| ------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------ | -------------------- | --------------------------------------------- |
| `downstreamBaseURL` | The base URL of the downstream API to forward traffic to                                                                                                                                                                             | :heavy_check_mark: | `-`                  | Config File or `DOWNSTREAM_BASE_URL` env var  |
| `SPEAKEASY_API_KEY` | The Speakeasy API Key (retrieved from the [Speakeasy Platform](https://app.speakeasyapi.dev/)) to use when communicating with Speakeasy                                                                                              | :heavy_check_mark: | `-`                  | `SPEAKEASY_API_KEY` env var only              |
| `apiID`             | The ID of the API to capture traffic for                                                                                                                                                                                             | :heavy_check_mark: | `-`                  | Config File or `SPEAKEASY_API_ID` env var     |
| `versionID`         | The Version of the API version to capture traffic for                                                                                                                                                                                | :heavy_check_mark: | `-`                  | Config File or `SPEAKEASY_VERSION_ID` env var |
| `openAPIDocs`       | A list of OpenAPI documents to use to associate captured traffic with the API. If multiple are provided they are merged with the [Speakeasy Merge Functionality](https://github.com/speakeasy-api/speakeasy/blob/main/docs/merge.md) | :heavy_minus_sign: | `["./openapi.yaml"]` | Config File or `OPENAPI_DOCS` env var         |
| `configLocation`    | The location of the configuration file to use                                                                                                                                                                                        | :heavy_minus_sign: | `./config.yaml`      | `CONFIG_LOCATION` env var only                |
| `port`              | The port to listen on                                                                                                                                                                                                                | :heavy_minus_sign: | `3333`               | Config File or `PORT` env var                 |

### Running the Proxy

The proxy can be built and run either as a binary or as a Docker image currently, allowing for a number of different deployment scenarios, such as running locally, in simple container based cloud environments or in more complex container orchestration environments like Kubernetes.

To run the proxy follow the steps below:

1. Create and obtain an API Key from the [Speakeasy Platform](https://app.speakeasyapi.dev/) either by creating a new Workspace or using an existing one.
2. Request access to the Speakeasy Request Capture Beta in our [Slack Community](https://join.slack.com/t/speakeasy-dev/shared_invite/zt-1eih279u9-uahunmIavQZJpiGmEIqYbA)
3. Create a configuration file (see [Configuration](#configuration) for details) ie. `./config.yaml` or set the required environment variables.
4. Create or use an existing OpenAPI document for the API you wish to capture traffic for and make sure it is accessible to the proxy ie. `./openapi.yaml`
5. Run the proxy either as a binary or as a Docker image (see [Running](#running) for details)
6. View the captured traffic in the [Speakeasy Platform](https://app.speakeasyapi.dev/) in your worksspace.

### Testing the Proxy locally

If you just want to test our the proxy locally, there are a few different ways you can do this:

1. The first is via our [Speakeasy CLI tool](https://github.com/speakeasy-api/speakeasy) (Documentation [here](https://github.com/speakeasy-api/speakeasy/blob/main/docs/proxy.md)) which allows running the proxy locally and capturing traffic for any HTTP (HTTPS terminatation not supported) API easily using simple command line options to configure the proxy.
   1. Request access to the Speakeasy Request Capture Beta in our [Slack Community](https://join.slack.com/t/speakeasy-dev/shared_invite/zt-1eih279u9-uahunmIavQZJpiGmEIqYbA)
   2. Authenticate the CLI tool using the `speakeasy auth login` command.
   3. Run the proxy using the Speakeasy CLI tool: `speakeasy proxy --api-id {API_ID} --version-id {VERSION_ID} --schema {OPENAPI_DOC_LOCATION} --downstream {DOWNSTREAM_BASE_URL}`
   4. Use the API you are capturing traffic for.
   5. View the captured traffic in the [Speakeasy Platform](https://app.speakeasyapi.dev/) in your workspace.
2. The second way is using the [httpbin](https://httpbin.org/) service as a test API to test it out, which we have setup to run using Docker Compose. Instruction for this are as follows:
   1. Follow steps 1-2 from [Running the Proxy](#running-the-proxy) above.
   2. Start the httpbin service and proxy using docker-compose: `SPEAKEASY_API_KEY='{API_KEY_HERE}' docker-compose up`
   3. Access the httpbin service via the proxy at <http://localhost:3333> and use some of the endpoints to generate traffic.
   4. View the captured traffic in the [Speakeasy Platform](https://app.speakeasyapi.dev/) in your workspace.

## Building

### Build Binary

```bash
go build -o speakeasy-proxy
```

### Build Docker Image

```bash
docker build -t speakeasy-proxy .
```

## Running

### Run Binary

```bash
SPEAKEASY_API_KEY='{API_KEY_HERE}' ./speakeasy-proxy
```

### Run Docker Image

If using our pre-built image first pull the required image:

```bash
docker pull ghcr.io/speakeasy-api/speakeasy-proxy:latest
```

otherwise build the image yourself (see [Building](#building) for details).

Then run the image:

```bash
docker run --network="host" -p 3333:3333 -v $(pwd)/config.yaml:/config.yaml -v $(pwd)/httpbin.yaml:/httpbin.yaml -e SPEAKEASY_API_KEY='{API_KEY_HERE}' speakeasy-proxy
