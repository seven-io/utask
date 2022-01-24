# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a uTask plugin for the Seven.io SMS service. It's built as a Go module that extends the OVH uTask framework to provide SMS sending capabilities through the Seven.io API.

## Architecture

- **Plugin Structure**: Built using the uTask plugin system (`github.com/ovh/utask/pkg/plugins/taskplugin`)
- **Main Plugin File**: `seven.go` - Contains the plugin implementation with configuration validation and execution logic
- **External Dependencies**:
  - `github.com/ovh/utask v1.29.1` - Core uTask framework
  - `github.com/seven-io/go-client v0.0.0-20231115115348-c0dea2fc0fe3` - Seven.io API client

## Key Components

- **Plugin Registration**: The plugin is registered with name "seven_utask" version "v0.1"
- **Configuration**: Expects `ApiKey` and `SmsParams` configuration, with `SEVEN_API_KEY` environment variable fallback
- **Execution**: Sends SMS via Seven.io API and returns JSON response or error

## Common Commands

```bash
# Build the module
go build

# Run tests
go test ./...

# Format code
go fmt ./...

# Vet code for issues
go vet ./...

# Check module dependencies
go mod verify

# Tidy dependencies
go mod tidy
```

## Development Notes

- Go version requirement: 1.19+
- The plugin follows uTask's plugin interface with `validConfig` and `exec` functions
- API key can be provided via configuration or `SEVEN_API_KEY` environment variable
- Error handling includes both SMS API errors and configuration validation errors