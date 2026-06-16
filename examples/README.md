# Examples

Each subdirectory contains a standalone runnable example for one API product.
The `all/` example demonstrates the aggregated `V1Client` (all products bundled).

## Running

```bash
# V1Client example (all products)
go run -C examples/all .

# Single product example
go run -C examples/sms .
go run -C examples/scoring .
```

## Configuration

Examples no longer contain hardcoded credentials. They call `var conf ksOpen.Config; conf.LoadEnv()` which reads configuration exclusively from environment variables:

| Variable | Config Field | Description |
|----------|-------------|-------------|
| `KS_CLIENT_ID` | `ClientID` | OAuth2 client ID |
| `KS_CLIENT_SECRET` | `ClientSecret` | OAuth2 client secret |
| `KS_SERVER_URL` | `ServerURL` | API gateway base URL (default: `https://api-gateway.kyivstar.ua`) |
| `KS_SERVER_MODE` | `ServerMode` | Server mode: `"mock"`, `"sandbox"`, or `"live"` |

```bash
export KS_CLIENT_ID=your_client_id
export KS_CLIENT_SECRET=your_client_secret
export KS_SERVER_URL=https://api-gateway.kyivstar.ua
go run -C examples/sms .
```

The same environment variables are used by the integration tests in `test/`.

## Config from JSON

You can also load config from a JSON file using `conf.LoadJSON(r io.Reader)`:

```go
f, _ := os.Open("config.json")
conf.LoadJSON(f)
```

JSON keys match the struct field names: `serverUrl`, `clientId`, `clientSecret`, `serverMode`.
