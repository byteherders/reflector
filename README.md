# reflector

[![Build Status](https://github.com/byteherders/reflector/actions/workflows/ci.yml/badge.svg)](https://github.com/byteherders/reflector/actions/workflows/ci.yml)
[![Latest Release](https://img.shields.io/github/v/release/byteherders/reflector?sort=semver)](https://github.com/byteherders/reflector/releases/latest)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/byteherders?label=Sponsor)](https://github.com/sponsors/byteherders)

Reflector is a tiny HTTP service that shows you *exactly* what an upstream sees from a request. It captures transport data, headers, cookies, query params, request body previews, TLS details, and—when loaded in a browser—rich metadata that the browser itself can collect. It is ideal for debugging CDN/CDN↔origin behavior, reverse-proxy issues, tracing weird ingress proxy/Kubernetes errors, and verifying what user-agents really send.

## Highlights

- 🔍 Mirrors every request attribute the Go `net/http` server receives, including `X-Forwarded-*` hints.
- 📄 Returns a Bootstrap-styled dashboard instead of raw JSON for quick, readable inspection.
- 🧠 Bundles a lightweight client-side collector that POSTs browser metadata (UA, storage, hardware, media permissions, etc.) back to the server automatically.
- 🩺 Ships with `/healthz` for uptime checks and configurable limits on how much of the body is captured.

## Quick start

```bash
git clone https://github.com/byteherder/reflector.git
cd reflector
go run ./cmd/reflector
```

By default the service listens on `:8080`. Visit `http://localhost:8080/` in a browser to see the rendered page. Open the same URL via `curl` or any HTTP client to inspect non-browser requests.

### Building a binary

```bash
go build -o reflector ./cmd/reflector
./reflector --port 9090 --body-bytes 8192
```

The binary honors the same flags:

| Flag | ENV | Description | Default |
| ---- | --- | ----------- | ------- |
| `--port` | `PORT` | TCP port to bind | `8080` |
| `--body-bytes` | – | Max number of request body bytes to capture | `4096` |

## Endpoints

| Path | Method | Purpose |
| ---- | ------ | ------- |
| `/` | GET/POST/etc. | Primary reflection page; automatically loads the browser collector script. |
| `/collect` | POST | Receives JSON metadata from the inline browser script (handled automatically). |
| `/healthz` | GET | Always returns `200 OK` for readiness/liveness probes. |

## Browser metadata collection

When you open `/` in a browser, Reflector injects a script that gathers:

- Navigator data (UA string, languages, hardware concurrency, do-not-track, etc.)
- Screen + viewport sizes, color depth, and preferred color scheme
- Network estimates (Connection API), performance memory, storage keys presence
- Touch/pointer support and presence of media devices

The script POSTs these details to `/collect`. The server re-renders the page to include a prettified JSON block beneath "Browser Metadata". This flow is automatic and requires no extra configuration.

## Deployment tips

- **Behind a CDN / proxy:** Ensure your proxy forwards `X-Forwarded-For`, `X-Forwarded-Proto`, and `X-Real-IP` if you rely on client IP visibility.
- **HTTPS/TLS:** Terminate TLS at your edge or wrap reflector with something like Caddy/Nginx; the TLS card will show the negotiated details if reflector terminates TLS itself.
- **Resource limits:** Use `--body-bytes` to avoid dumping large payloads into the response; set it to `0` if you want to disable body capture entirely.

## Development

```bash
# Run tests/build in an isolated cache
GOCACHE=$(pwd)/.cache/go-build \
GOMODCACHE=$(pwd)/.cache/pkg/mod \
GOPATH=$(pwd)/.cache \
go test ./...
```

The core server lives under [`internal/server`](internal/server) and the CLI entry point is in [`cmd/reflector`](cmd/reflector). Contributions and issue reports are welcome via GitHub.

## License

Apache-2.0. See [LICENSE](LICENSE) for details.
