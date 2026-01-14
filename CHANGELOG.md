# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- OpenTelemetry tracing support via `NewClientWithOTEL` and `NewClientWithResponsesWithOTEL`

## [1.1.0] - 2026-01-14

### Changed

- Update minimum Go version to 1.24 for broader compatibility
- Add Go 1.24/1.25 build matrix to CI workflow
- Bump actions/checkout from v4 to v6
- Bump actions/setup-go from v5 to v6

## [1.0.0] - 2026-01-14

### Added

- Initial Go client generated from OpenAPI spec
- Authentication endpoint (`/api/v1/auth/authenticate`)
- Create/Update shipment endpoint (`/api/v1/shipment/create`)
- Create shipment invoice endpoint (`/api/v1/shipment/issueInvoice`)
- GitHub Actions CI workflow (build, lint, generated code check)
- GitHub Actions release workflow
- Dependabot configuration
