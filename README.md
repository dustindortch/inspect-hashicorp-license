# Inspect Hashicorp License

CLI tool to read and inspect Hashicorp licenses.

## Usage

```shell
licinstect -file <path-to-license-file>
```

output:

```json
{
  "license_id": "aaaaaaaa-bbbb-cccc-dddd-000000000000",
  "customer_id": "aaaaaaaa-bbbb-cccc-dddd-000000000000",
  "installation_id": "*",
  "issue_time": "1970-01-01T00:00:00Z",
  "start_time": "1970-01-01T00:00:00Z",
  "expiration_time": "1970-01-01T00:00:00.00Z",
  "termination_time": "1970-01-01T00:00:00.00Z",
  "product": "vault",
  "flags": {
    "modules": [
      "multi-dc-scale",
      "governance-policy",
      "advanced-data-protection-transform",
      "advanced-data-protection-key-management"
    ]
  }
}
```

## Build

```shell
go build -o licinspect .
```
