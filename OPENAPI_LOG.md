# OpenAPI Spec Inconsistencies Log

This document tracks inconsistencies between the upstream EPGL API OpenAPI specification and the actual API behavior. These fixes are applied to `spec/apispec.json` to ensure the generated Go client works correctly.

---

## 2026-01-14: Error Response Schema Mismatch

**Issue:** [#5](https://github.com/quiqupltd/epglgo/issues/5)

**Problem:**
The original OpenAPI spec defined the 400 error response `errors` field as:
```yaml
errors:
  type: object
  additionalProperties:
    type: array
    items:
      type: string
```

This generated `Errors *map[string][]string` in Go, but the actual API returns:
```json
{
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
  "timestamp": "2026-01-14T08:57:44.7958663Z"
}
```

**Impact:**
JSON unmarshal error: `cannot unmarshal array into Go struct field .errors of type map[string][]string`

**Fix Applied:**
Created `ErrorResponse` schema in `components/schemas` with correct structure:
- `status` (string)
- `message` (string)
- `data` (nullable)
- `errors` (array of objects with `code`, `message`, `field`)
- `correlationId` (string, uuid)
- `timestamp` (string, date-time)

**Affected Endpoints:**
- `POST /api/v1/shipment/create` (400 response)
- `POST /api/v1/shipment/issueInvoice` (400 response)
