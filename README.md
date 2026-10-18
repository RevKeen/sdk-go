# RevKeen Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/revkeen/sdk-go.svg)](https://pkg.go.dev/github.com/revkeen/sdk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Official Go client for the [RevKeen API](https://docs.revkeen.com/api-reference/openapi) — auto-generated from the OpenAPI specification via [OpenAPI Generator](https://openapi-generator.tech).

> [!IMPORTANT]
> This repository is a generated release mirror. Propose source changes in
> [RevKeen/revkeen](https://github.com/RevKeen/revkeen); `main` here is updated only by an explicit reviewed SDK release.

## Installation

```bash
go get github.com/revkeen/sdk-go
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

    revkeen "github.com/revkeen/sdk-go"
)

func main() {
    client, err := revkeen.NewClient(
        os.Getenv("REVKEEN_API_KEY"),
        revkeen.EnvironmentSandbox,
    )
    if err != nil {
        log.Fatal(err)
    }

    customers, _, err := client.CustomersAPI.
        CustomersList(context.Background()).
        Limit(10).
        Execute()
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
client, err := revkeen.NewClient(
    os.Getenv("REVKEEN_API_KEY"),
    revkeen.EnvironmentProduction,
)
```

## Resources

OpenAPI Generator exposes typed service fields such as `CustomersAPI`,
`InvoicesAPI`, and `SubscriptionsAPI`. Each operation uses the generated
request builder and ends with `Execute()`. Refer to the generated API docs for
the current method list.

## Error Handling

```go
_, _, err := client.CustomersAPI.CustomersGet(context.Background(), "cus_nonexistent").Execute()
if err != nil {
    var apiErr *revkeen.RequestError
    if errors.As(err, &apiErr) {
        fmt.Printf("API error %d (%s), request %s\n", apiErr.StatusCode, apiErr.Code, apiErr.RequestID)
    }
}
```

## Configuration

```go
client, err := revkeen.NewClient(
    os.Getenv("REVKEEN_API_KEY"),
    revkeen.EnvironmentSandbox,
)

// Custom endpoints require the explicit escape hatch.
local, err := revkeen.NewClientWithCustomBaseURL(
    os.Getenv("REVKEEN_API_KEY"),
    "http://localhost:4010/v2",
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
