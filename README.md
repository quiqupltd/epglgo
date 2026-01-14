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

## Development

### Regenerate client from OpenAPI spec

```bash
go generate ./...
```

### OpenAPI spec location

The OpenAPI specification is at `spec/apispec.json`.
