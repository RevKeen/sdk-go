# RevKeen Go SDK

The official Go client for the [RevKeen API](https://docs.revkeen.com). Idiomatic `context.Context` plumbing, typed request and response structs, auto-pagination iterators, webhook verification, and production-grade retries.

[![Go Reference](https://pkg.go.dev/badge/github.com/RevKeen/sdk-go.svg)](https://pkg.go.dev/github.com/RevKeen/sdk-go)
[![CI](https://img.shields.io/github/actions/workflow/status/RevKeen/sdk-go/ci.yml?branch=main&style=flat-square&label=ci)](https://github.com/RevKeen/sdk-go/actions)
[![License: MIT](https://img.shields.io/badge/license-MIT-000?style=flat-square)](./LICENSE)
[![Docs](https://img.shields.io/badge/docs-docs.revkeen.com-000?style=flat-square)](https://docs.revkeen.com/docs/sdks/go)

## Install

```bash
go get github.com/RevKeen/sdk-go
```

Requires Go 1.21+.

## Quick start

```go
package main

import (
    "context"
    "fmt"
    "os"

    revkeen "github.com/RevKeen/sdk-go"
)

func main() {
    client := revkeen.NewClient(revkeen.Config{
        APIKey: os.Getenv("REVKEEN_API_KEY"),
    })

    customer, err := client.Customers.Create(context.Background(), &revkeen.CustomerCreateParams{
        Email: revkeen.String("ops@acme.example"),
        Name:  revkeen.String("Acme Inc."),
    })
    if err != nil {
        panic(err)
    }

    fmt.Println(customer.ID)
}
```

## Features

- **Typed request and response structs** — no `map[string]interface{}` anywhere
- **Context-aware** — every method takes `context.Context` for cancellation and deadlines
- **Automatic pagination iterators** — `iter.Next()` / `iter.Current()` / `iter.Err()`
- **Automatic retries** — exponential backoff on `5xx`, `429`, network errors
- **Idempotency keys** — attached automatically on safe-to-retry mutations
- **Webhook verification** — `revkeen.Webhooks.Verify(body, signature, secret)`
- **OAuth 2.1 + API-key auth** — both first-class

## Documentation

- [SDK docs](https://docs.revkeen.com/docs/sdks/go) — examples, recipes, and full API surface
- [API reference](https://docs.revkeen.com/docs/api-reference) — every endpoint, from the OpenAPI spec
- [pkg.go.dev](https://pkg.go.dev/github.com/RevKeen/sdk-go) — auto-generated godoc reference
- [Webhooks guide](https://docs.revkeen.com/docs/webhooks) — signature verification + event catalogue
- [Versioning](https://docs.revkeen.com/docs/fundamentals/versioning) — API ↔ SDK compatibility matrix

## Generation

This SDK is generated from the [canonical OpenAPI spec](https://docs.revkeen.com/docs/api-reference). The generator runs on every spec change. A human-authored layer adds idiomatic helpers for pagination, retries, webhooks, and errors.

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for development setup, test instructions, and the release process.

Please file issues and feature requests on the [issue tracker](https://github.com/RevKeen/sdk-go/issues). For security disclosures, see [SECURITY.md](./SECURITY.md).

## License

[MIT](./LICENSE) — © RevKeen.
