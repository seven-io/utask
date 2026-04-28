<p align="center">
  <img src="https://www.seven.io/wp-content/uploads/Logo.svg" width="250" alt="seven logo" />
</p>

<h1 align="center">seven SMS for uTask</h1>

<p align="center">
  Plugin for OVH's <a href="https://github.com/ovh/utask">uTask</a> workflow engine that sends SMS via the seven gateway as a regular workflow step.
</p>

<p align="center">
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-teal.svg" alt="MIT License" /></a>
  <img src="https://img.shields.io/badge/uTask-1.29%2B-blue" alt="uTask 1.29+" />
  <img src="https://img.shields.io/badge/Go-1.19%2B-00add8" alt="Go 1.19+" />
</p>

---

## Features

- **Native uTask Plugin** - Registered as `seven_utask` (v0.1)
- **Workflow Step** - Drop into any uTask YAML workflow as a regular step
- **Env or Config Auth** - API key via `SEVEN_API_KEY` env or per-task config

## Prerequisites

- Go 1.19+
- A uTask installation (>= v1.29.1)
- A [seven account](https://www.seven.io/) with API key ([How to get your API key](https://help.seven.io/en/developer/where-do-i-find-my-api-key))

## Installation

```bash
git clone git@github.com:seven-io/utask.git
cd utask
go build
```

Drop the resulting binary or built plugin into your uTask plugin path.

## Configuration

### Option A: Environment variable

```bash
export SEVEN_API_KEY=your-seven-api-key
```

### Option B: Per-task config

```yaml
config:
  apiKey: your-seven-api-key
  smsParams:
    to:   "+1234567890"
    text: "Hello from uTask!"
```

## Usage

Use the plugin as any other workflow step:

```yaml
steps:
  - name: send-sms
    plugin: seven_utask
    config:
      smsParams:
        to:   "+1234567890"
        text: "Your workflow notification message"
```

### `smsParams`

`smsParams` accepts every parameter supported by the seven SMS API:

| Field | Description |
|-------|-------------|
| `to` | Recipient phone number (required) |
| `text` | Message body (required) |
| `from` | Sender ID. Up to 11 alphanumeric or 16 numeric characters |
| `delay` | Delayed dispatch as Unix timestamp |
| ... | All other [seven SMS API parameters](https://docs.seven.io/en/rest-api/endpoints/sms#send-sms) |

## Development

```bash
go fmt ./...
go vet ./...
go test ./...
go mod tidy
```

## Support

Need help? Feel free to [contact us](https://www.seven.io/en/company/contact/) or [open an issue](https://github.com/seven-io/utask/issues).

## License

[MIT](LICENSE)
