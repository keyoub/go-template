<!--
Sync Impact Report:
- Version change: none → 1.0.0
- Added principles: I. Clean Architecture for Go Web Services, II. Data Contract Isolation, III. Comprehensive Testing Discipline, IV. No ORMs and Standard Library Preference, V. Integration Testing and Task Management
- Added sections: Documentation Standards, Development Workflow
- Removed sections: none
- Templates requiring updates: none
- Follow-up TODOs: none
-->
# Go Web App Template Constitution

## Core Principles

### I. Clean Architecture for Go Web Services
Follow the clean architecture principles outlined in https://www-all.easypost.com/blog/2022-10-24-clean-go-for-web-services. Services MUST be structured into four distinct segments: Repository packages providing data access interfaces, Entity package implementing shared data constructs, Service packages containing application/business logic, and External interface packages (HTTP controllers or CLI commands). The Entity package MUST hold contracts between repo, service, api, and cmd packages, allowing external interfaces to scale independently of intra-service dependencies. This layered approach ensures dependency inversion, testability, and clean migrations of integrations.

### II. Data Contract Isolation
Data contracts used between packages MUST NOT have json/db tags. These MUST be isolated to the struct/contract of the communication interface, whether responding to HTTP APIs or sending/receiving data from db. Internal packages MUST use clean structs without serialization tags to maintain separation of concerns and prevent tight coupling to external formats.

### III. Comprehensive Testing Discipline
Every func MUST have a unit test that covers all scenarios, including error paths for the specific func. Dependencies/integrations MUST be mocked to ensure unit test is isolated to just the func being tested. There MUST always be a mock of the repo interfaces for unit testing in a shared testutils pkg. This ensures high test coverage, reliability, and fast feedback during development.

### IV. No ORMs and Standard Library Preference
No ORMs are ever used. The app MUST use as much of the standard library as possible. Only bring in 3rd party pkgs when absolutely necessary (e.g. for routing, logging, etc). All 3rd party pkgs MUST be well maintained and widely used in the go community. They MUST be vetted for license compatibility as well as security vulnerabilities. This minimizes dependencies, reduces attack surface, and ensures long-term maintainability.

### V. Integration Testing and Task Management
There MUST be a clear integrations src/app/tests pkg that has standalone integration tests that spins up the whole app as a black box and tests all the real dependencies (db, external APIs, etc) to ensure the app works as expected in a real world scenario. Justfile MUST be used to manage common tasks like running tests, building the app, linting, etc. Integration tests MUST verify end-to-end functionality, while Justfile provides consistent, reproducible build and test workflows.

## Documentation Standards

The web app MUST have a well documented openapi spec in the top level ./openapi_specs/ directory that outlines all the HTTP APIs exposed by the app. This spec MUST be kept up to date as the app evolves. The web app MUST have a clear ARCHITECTURE.md doc that outlines the high level architecture of the app that includes mermaid diagrams and request flows. It MUST document each distinct feature of the app along its architecture. The web app MUST have a clear README.md doc that outlines how to get started with the app, how to run it locally, how to run tests, how to build/deploy it, and any other relevant information for developers working on the app.

## Development Workflow

Code MUST follow the clean architecture layers with clear package boundaries. Changes to data stores or external integrations MUST be isolated to repo packages without affecting service or api layers. Unit tests MUST use mocks from testutils pkg. Integration tests MUST run against real dependencies. Justfile MUST define all common development tasks. All 3rd party dependencies MUST be reviewed for security and license before inclusion.

## Governance

This constitution supersedes all other practices for the Go web app template. Amendments require documentation of the change rationale, approval from maintainers, and a migration plan for existing projects using the template. All code contributions MUST verify compliance with these principles. Complexity MUST be justified with clear benefits. Use this constitution as the primary guidance for development decisions.

**Version**: 1.0.0 | **Ratified**: 2025-10-22 | **Last Amended**: 2025-10-22
