# speakeasy-proxy

## Build Docker Image

```bash
docker build -t speakeasy-proxy .
```

## Run Docker Image

```bash
docker run --network="host" -p 3333:3333 -v $(pwd)/config.yaml:/config.yaml -v $(pwd)/httpbin.yaml:/httpbin.yaml -e SPEAKEASY_API_KEY='{API_KEY_HERE}' speakeasy-proxy
```