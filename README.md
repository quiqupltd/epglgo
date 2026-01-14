# epglgo

Go client for the EPGL (Regulatory and Licensing) API.

## Installation

```bash
go get github.com/quiqupltd/epglgo
```

## Usage

### Create a client

```go
client, err := epglgo.NewClient("https://api-stg.epgl.ae")
if err != nil {
    log.Fatal(err)
}
```

### Authenticate

```go
resp, err := client.AuthenticateClient(ctx, epglgo.AuthenticateClientJSONRequestBody{
    ClientId:     "your-client-id",
    ClientSecret: "your-client-secret",
})
```

### Create/Update a shipment

```go
resp, err := client.CreateUpdateShipment(ctx, epglgo.CreateUpdateShipmentJSONRequestBody{
    // ... shipment details
})
```

### Create a shipment invoice

```go
resp, err := client.CreateShipmentInvoice(ctx, epglgo.CreateShipmentInvoiceJSONRequestBody{
    // ... invoice details
})
```

### Using Bearer authentication

```go
client, err := epglgo.NewClient("https://api-stg.epgl.ae",
    epglgo.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
        req.Header.Set("Authorization", "Bearer "+token)
        return nil
    }),
)
```

### OpenTelemetry Tracing

The client supports OpenTelemetry tracing out of the box. Use the OTEL-enabled constructors to automatically instrument all HTTP requests with distributed tracing:

```go
// Create a client with OpenTelemetry tracing
client, err := epglgo.NewClientWithResponsesWithOTEL("https://api-stg.epgl.ae")
if err != nil {
    log.Fatal(err)
}

// All requests will automatically create spans
resp, err := client.AuthenticateClientWithResponse(ctx, epglgo.AuthenticateClientJSONRequestBody{
    ClientId:     "your-client-id",
    ClientSecret: "your-client-secret",
})
```

You can combine OTEL tracing with other client options:

```go
client, err := epglgo.NewClientWithResponsesWithOTEL("https://api-stg.epgl.ae",
    epglgo.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
        req.Header.Set("Authorization", "Bearer "+token)
        return nil
    }),
)
```

**Note:** You must configure an OpenTelemetry trace provider in your application for traces to be exported. See the [OpenTelemetry Go documentation](https://opentelemetry.io/docs/languages/go/) for setup instructions.

## Development

### Regenerate client from OpenAPI spec

```bash
go generate ./...
```

### OpenAPI spec location

The OpenAPI specification is at `spec/apispec.json`.
