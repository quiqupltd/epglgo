package epglgo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("https://api-stg.epgl.ae")
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	if client == nil {
		t.Fatal("NewClient() returned nil client")
	}
	// Server URL may have trailing slash added
	if client.Server != "https://api-stg.epgl.ae" && client.Server != "https://api-stg.epgl.ae/" {
		t.Errorf("NewClient() server = %v, want %v", client.Server, "https://api-stg.epgl.ae")
	}
}

func TestNewClientWithResponses(t *testing.T) {
	client, err := NewClientWithResponses("https://api-stg.epgl.ae")
	if err != nil {
		t.Fatalf("NewClientWithResponses() error = %v", err)
	}
	if client == nil {
		t.Fatal("NewClientWithResponses() returned nil client")
	}
}

func TestNewClientWithOTEL(t *testing.T) {
	client, err := NewClientWithOTEL("https://api-stg.epgl.ae")
	if err != nil {
		t.Fatalf("NewClientWithOTEL() error = %v", err)
	}
	if client == nil {
		t.Fatal("NewClientWithOTEL() returned nil client")
	}
}

func TestNewClientWithResponsesWithOTEL(t *testing.T) {
	client, err := NewClientWithResponsesWithOTEL("https://api-stg.epgl.ae")
	if err != nil {
		t.Fatalf("NewClientWithResponsesWithOTEL() error = %v", err)
	}
	if client == nil {
		t.Fatal("NewClientWithResponsesWithOTEL() returned nil client")
	}
}

func TestWithRequestEditorFn(t *testing.T) {
	editor := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("X-Custom-Header", "test-value")
		return nil
	}

	client, err := NewClient("https://api-stg.epgl.ae", WithRequestEditorFn(editor))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	if len(client.RequestEditors) != 1 {
		t.Errorf("Expected 1 request editor, got %d", len(client.RequestEditors))
	}
}

func TestWithHTTPClient(t *testing.T) {
	customClient := &http.Client{}
	client, err := NewClient("https://api-stg.epgl.ae", WithHTTPClient(customClient))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	if client.Client != customClient {
		t.Error("WithHTTPClient() did not set custom HTTP client")
	}
}

func TestErrorResponseUnmarshal(t *testing.T) {
	// Test that ErrorResponse correctly unmarshals the actual API error format
	jsonData := `{
		"status": "Error",
		"message": "Validation failed.",
		"data": null,
		"errors": [
			{
				"code": "VALIDATION_ERROR",
				"message": "unit is required.",
				"field": "details.declaredWeight.unit"
			}
		],
		"correlationId": "464ddc5e-b9ad-4560-a11d-b79464ebb688",
		"timestamp": "2026-01-14T08:57:44.795866Z"
	}`

	var errResp ErrorResponse
	err := json.Unmarshal([]byte(jsonData), &errResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal ErrorResponse: %v", err)
	}

	if errResp.Status == nil || *errResp.Status != "Error" {
		t.Errorf("Expected status 'Error', got %v", errResp.Status)
	}
	if errResp.Message == nil || *errResp.Message != "Validation failed." {
		t.Errorf("Expected message 'Validation failed.', got %v", errResp.Message)
	}
	if errResp.Errors == nil || len(*errResp.Errors) != 1 {
		t.Fatalf("Expected 1 error, got %v", errResp.Errors)
	}

	firstError := (*errResp.Errors)[0]
	if firstError.Code == nil || *firstError.Code != "VALIDATION_ERROR" {
		t.Errorf("Expected error code 'VALIDATION_ERROR', got %v", firstError.Code)
	}
	if firstError.Field == nil || *firstError.Field != "details.declaredWeight.unit" {
		t.Errorf("Expected field 'details.declaredWeight.unit', got %v", firstError.Field)
	}
	if firstError.Message == nil || *firstError.Message != "unit is required." {
		t.Errorf("Expected error message 'unit is required.', got %v", firstError.Message)
	}
}

func TestAuthenticateClientRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v1/auth/authenticate" {
			t.Errorf("Expected path /api/v1/auth/authenticate, got %s", r.URL.Path)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Return mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"token": "mock-jwt-token"}`))
	}))
	defer server.Close()

	client, err := NewClientWithResponses(server.URL)
	if err != nil {
		t.Fatalf("NewClientWithResponses() error = %v", err)
	}

	resp, err := client.AuthenticateClientWithResponse(context.Background(), AuthenticateClientJSONRequestBody{
		ClientId:     "test-client-id",
		ClientSecret: "test-client-secret",
	})
	if err != nil {
		t.Fatalf("AuthenticateClientWithResponse() error = %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode())
	}
}

func TestCreateUpdateShipmentRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v1/shipment/create" {
			t.Errorf("Expected path /api/v1/shipment/create, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success"}`))
	}))
	defer server.Close()

	client, err := NewClient(server.URL)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	resp, err := client.CreateUpdateShipment(context.Background(), CreateUpdateShipmentJSONRequestBody{})
	if err != nil {
		t.Fatalf("CreateUpdateShipment() error = %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestCreateShipmentInvoiceRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v1/shipment/issueInvoice" {
			t.Errorf("Expected path /api/v1/shipment/issueInvoice, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success"}`))
	}))
	defer server.Close()

	client, err := NewClient(server.URL)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	resp, err := client.CreateShipmentInvoice(context.Background(), CreateShipmentInvoiceJSONRequestBody{})
	if err != nil {
		t.Fatalf("CreateShipmentInvoice() error = %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
