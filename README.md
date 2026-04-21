# RevKeen Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/revkeen/revkeen-go.svg)](https://pkg.go.dev/github.com/revkeen/revkeen-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Official Go client for the [RevKeen API](https://docs.revkeen.com/api-reference/openapi) — auto-generated from the OpenAPI specification via [OpenAPI Generator](https://openapi-generator.tech).

## Installation

```bash
go get github.com/revkeen/revkeen-go
```

Requires Go 1.22 or later.

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    revkeen "github.com/revkeen/revkeen-go"
    "github.com/revkeen/revkeen-go/option"
)

func main() {
    client := revkeen.NewClient(
        option.WithAPIKey(os.Getenv("REVKEEN_API_KEY")),
    )

    customers, err := client.Customers.List(context.Background(), &revkeen.CustomersListRequest{
        Limit: revkeen.Int(10),
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, c := range customers.Data {
        fmt.Println(c.Name, c.Email)
    }
}
```

## Authentication

### API Key (recommended for server-to-server)

```go
client := revkeen.NewClient(
    option.WithAPIKey(os.Getenv("REVKEEN_API_KEY")),
)
```

### OAuth 2.1 (recommended for MCP and third-party integrations)

```go
client := revkeen.NewClient(
    option.WithOAuth(revkeen.OAuthConfig{
        ClientID:     os.Getenv("REVKEEN_CLIENT_ID"),
        ClientSecret: os.Getenv("REVKEEN_CLIENT_SECRET"),
        Scopes:       []string{"customers:read", "invoices:read"},
    }),
)
```

See the [OAuth guide](https://docs.revkeen.com/docs/developers/oauth) for details.

## Resources

Every API resource is available as a typed field on the client:

| Resource | Method examples |
|----------|----------------|
| `client.Customers` | `List()`, `Create()`, `Get()`, `Update()`, `Delete()` |
| `client.Invoices` | `List()`, `Create()`, `Get()`, `Update()`, `Finalize()`, `Send()`, `Void()` |
| `client.Subscriptions` | `List()`, `Create()`, `Get()`, `Update()`, `Cancel()`, `Pause()`, `Resume()` |
| `client.Products` | `List()`, `Create()`, `Get()`, `Update()`, `Delete()` |
| `client.Payments` | `List()`, `Create()`, `Get()` |
| `client.CheckoutSessions` | `Create()`, `Get()` |
| `client.Discounts` | `List()`, `Create()`, `Get()`, `Update()`, `Delete()` |
| `client.CreditNotes` | `List()`, `Create()`, `Get()` |
| `client.PaymentLinks` | `List()`, `Create()`, `Get()`, `Update()` |
| `client.PaymentMethods` | `List()`, `Get()`, `Detach()` |
| `client.WebhookEndpoints` | `List()`, `Create()`, `Delete()` |
| `client.Events` | `List()`, `Get()` |
| `client.Entitlements` | `List()`, `Check()` |

## Error Handling

```go
_, err := client.Customers.Get(context.Background(), "cus_nonexistent")
if err != nil {
    var apiErr *revkeen.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API error %d: %s\n", apiErr.StatusCode, apiErr.Message)
    }
}
```

## Configuration

```go
client := revkeen.NewClient(
    option.WithAPIKey(os.Getenv("REVKEEN_API_KEY")),
    // Staging environment
    option.WithBaseURL("https://staging-api.revkeen.com"),
)
```

## Compatibility

- **Runtime:** Go 1.22+
- **Dependencies:** Standard library only

## Links

- [API Reference](https://docs.revkeen.com/api-reference/openapi)
- [SDK Documentation](https://docs.revkeen.com/docs/developers/sdks/go)
- [TypeScript SDK](https://github.com/revkeen/sdk-typescript)
- [PHP SDK](https://github.com/revkeen/sdk-php)
