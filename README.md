# Seven.io uTask Plugin

A uTask plugin that enables SMS sending through the [Seven.io](https://seven.io) API.

## Overview

This plugin extends the [OVH uTask](https://github.com/ovh/utask) workflow engine to provide SMS functionality using Seven.io's messaging service. It allows you to send SMS messages as part of your automated workflows.

## Installation

1. Clone this repository:
```bash
git clone https://github.com/seven-io/utask.git
cd utask
```

2. Build the plugin:
```bash
go build
```

## Configuration

The plugin requires a Seven.io API key to function. You can provide this in two ways:

### Option 1: Environment Variable
Set the `SEVEN_API_KEY` environment variable:
```bash
export SEVEN_API_KEY="your-api-key-here"
```

### Option 2: Task Configuration
Include the API key directly in your task configuration:
```yaml
config:
  apiKey: "your-api-key-here"
  smsParams:
    to: "+1234567890"
    text: "Hello from uTask!"
```

## Usage

Add the plugin to your uTask workflow:

```yaml
steps:
  - name: send-sms
    plugin: seven_utask
    config:
      smsParams:
        to: "+1234567890"
        text: "Your workflow notification message"
```

### SMS Parameters

The `smsParams` field accepts all parameters supported by the Seven.io SMS API:

- `to`: Recipient phone number (required)
- `text`: Message content (required)
- `from`: Sender ID (optional)
- `delay`: Delayed sending timestamp (optional)
- Additional parameters as supported by the Seven.io API

## API Key

To get your Seven.io API key:

1. Sign up at [seven.io](https://seven.io)
2. Navigate to your dashboard
3. Copy your API key from the account settings

## Requirements

- Go 1.19 or higher
- Valid Seven.io API account and key
- uTask framework v1.29.1 or compatible version

## Development

```bash
# Format code
go fmt ./...

# Run tests
go test ./...

# Check for issues
go vet ./...

# Update dependencies
go mod tidy
```

## License

This project follows the same license as the uTask framework.

## Support

For Seven.io API questions, visit [Seven.io Documentation](https://www.seven.io/en/docs/)
For uTask framework questions, visit [uTask Documentation](https://github.com/ovh/utask)