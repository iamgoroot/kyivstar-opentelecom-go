# Testing

The test suite in `test/` supports two modes: **local** (default) and **real server**.

## Local mode (no credentials)

```bash
go test ./test/... -v
```

Spawns an in-process `httptest.Server` with mock handlers. No network calls, no credentials needed. All tests in `test/` run, including error-response tests and promo multipart upload tests. Pool/buffer tests live in `internal/client` and run with `go test ./internal/client ./test/...` or `go test ./...`.

## Real server mode (with credentials)

Set four environment variables:

| Variable | Required | Description |
|----------|----------|-------------|
| `KS_CLIENT_ID` | yes | OAuth2 client ID |
| `KS_CLIENT_SECRET` | yes | OAuth2 client secret |
| `KS_SERVER_URL` | yes | API gateway URL (e.g. `https://api-gateway.kyivstar.ua`) |
| `KS_SERVER_MODE` | no | `mock`, `sandbox`, or `live` (default: `live`) |

```bash
KS_CLIENT_ID=abc KS_CLIENT_SECRET=xyz KS_SERVER_URL=https://api-gateway.kyivstar.ua KS_SERVER_MODE=mock go test ./test/... -v
```

The `setupTestClient` helper in `test/helpers_test.go` detects the env vars and creates an OAuth2-authenticated client against the real server. The `KS_SERVER_MODE` value is appended to the URL path (`/mock`, `/sandbox`, or nothing for live).

## Tests that only run locally

Only pool tests are skipped when real server env vars are set:

- **Pool tests** (`TestBufferPoolUnderMemoryPressure`, `TestConcurrentPoolRequests`) — buffer pool memory/corruption tests under GC pressure

These call `isRunningLocally()` and `t.Skip()` when a real server URL is configured.
