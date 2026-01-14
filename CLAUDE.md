# CLAUDE.md

## Project Overview

Go client library for the EPGL Regulatory and Licensing API, generated from OpenAPI spec using oapi-codegen.

## Git Workflow

This project uses **git-flow** for development:

- `main` - Production releases only
- `develop` - Integration branch for features
- `feature/*` - New features (branch from `develop`)
- `release/*` - Release preparation
- `hotfix/*` - Production hotfixes (branch from `main`)

## Project Structure

```
.
├── api.gen.go           # Generated client code (DO NOT EDIT)
├── otel.go              # Manual: OpenTelemetry tracing helpers
├── generate.go          # go:generate directive
├── .oapi-codegen.yml    # oapi-codegen configuration
├── OPENAPI_LOG.md       # Log of upstream API spec inconsistencies
├── spec/
│   └── apispec.json     # OpenAPI 3.0 specification (with fixes applied)
└── tmp/                 # Source API spec files (can be deleted after merge)
```

### Generated vs Manual Code

**Generated (DO NOT EDIT):**
- `api.gen.go` - Client, models, and request/response types from OpenAPI spec

**Manual (safe to edit):**
- `otel.go` - OpenTelemetry tracing wrappers
- Any other `*.go` files (except `api.gen.go`)

When adding new features (like OTEL), always create separate files. Never modify `api.gen.go` directly as changes will be lost on regeneration.

## Code Generation

Regenerate client after spec changes:

```bash
go generate ./...
```

The generator uses `oapi-codegen` configured in `.oapi-codegen.yml`.

## API Endpoints

- `POST /api/v1/auth/authenticate` - Authenticate and get JWT token
- `POST /api/v1/shipment/create` - Create or update a shipment
- `POST /api/v1/shipment/issueInvoice` - Create invoice for a shipment

## OpenAPI Spec Fixes

The upstream EPGL API spec sometimes has inconsistencies with actual API behavior. When fixing these:

1. Apply fixes to `spec/apispec.json`
2. Document the inconsistency in `OPENAPI_LOG.md` with:
   - Date and related issue
   - What the spec said vs actual API behavior
   - Impact on generated code
   - Fix applied
3. Regenerate the client with `go generate ./...`

## Important Notes

- `api.gen.go` is auto-generated - do not edit manually
- OpenAPI spec uses `number` type for decimals (mapped to `float64` in Go)
- Staging server: `https://api-stg.epgl.ae`
