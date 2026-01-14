package epglgo

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// NewClientWithOTEL creates a new Client with OpenTelemetry tracing enabled.
// The client will automatically create spans for all HTTP requests.
//
// Example usage:
//
//	client, err := epglgo.NewClientWithOTEL("https://api-stg.epgl.ae")
//	if err != nil {
//	    log.Fatal(err)
//	}
func NewClientWithOTEL(server string, opts ...ClientOption) (*Client, error) {
	opts = append([]ClientOption{WithHTTPClient(newOTELHTTPClient())}, opts...)
	return NewClient(server, opts...)
}

// NewClientWithResponsesWithOTEL creates a new ClientWithResponses with OpenTelemetry tracing enabled.
// The client will automatically create spans for all HTTP requests and parse responses.
//
// Example usage:
//
//	client, err := epglgo.NewClientWithResponsesWithOTEL("https://api-stg.epgl.ae")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	resp, err := client.AuthenticateClientWithResponse(ctx, epglgo.AuthenticateClientJSONRequestBody{
//	    ClientId:     "my-client-id",
//	    ClientSecret: "my-client-secret",
//	})
func NewClientWithResponsesWithOTEL(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	opts = append([]ClientOption{WithHTTPClient(newOTELHTTPClient())}, opts...)
	return NewClientWithResponses(server, opts...)
}

// newOTELHTTPClient creates an HTTP client with OpenTelemetry transport instrumentation.
func newOTELHTTPClient() *http.Client {
	return &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
}
