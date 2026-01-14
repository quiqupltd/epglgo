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
├── generate.go          # go:generate directive
├── .oapi-codegen.yml    # oapi-codegen configuration
├── spec/
│   └── apispec.json     # OpenAPI 3.0 specification
└── tmp/                 # Source API spec files (can be deleted after merge)
```

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

## Important Notes

- `api.gen.go` is auto-generated - do not edit manually
- OpenAPI spec uses `number` type for decimals (mapped to `float64` in Go)
- Staging server: `https://api-stg.epgl.ae`
