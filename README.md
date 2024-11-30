# SOPS Terraform Provider

A Terraform provider for [sops](https://getsops.io/) providing functions to decrypt SOPS-encrypted files and strings. This provider is heavily inspired by [`carlpett/terraform-provider-sops`](https://github.com/carlpett/terraform-provider-sops) but does not rely on data sources and instead uses provider functions. This allows you to decrypt secrets without them necessarily being stored in the Terraform state as plain text.

## Requirements

As provider functions are a fairly new feature in Terraform, you will need to be using Terraform v1.8 or later.

## Usage

This provider contains two provider functions:

- `file` - Decrypts a local file using SOPS
- `string` - Decrypts a string using SOPS, useful if the secret is not stored in a local file

To make use of the provider, you will need to add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    sops = {
      source = "nobbs/sops"
      version = "~> 0.1.0"
    }
  }
}

# The provider has no configuration options
provider "sops" {}
```

## Examples

See the [examples](./examples) directory for more examples.

Once the provider is configured, you can use the provider functions in your Terraform configuration as follows:

```hcl
output "basic-json" {
  value = provider::sops::file("./test/fixtures/basic.sops.json")
}

# Output:
# basic-json = {
#   "data" = {
#     "abc" = "xyz"
#     "floats" = 0.000000000314
#     "integers" = 123
#     "truthy" = true
#   }
#   "raw" = <<-EOT
#   {
#   	"abc": "xyz",
#   	"integers": 123,
#   	"truthy": true,
#   	"floats": 3.14e-10
#   }
#   EOT
# }
```

Or using the `string` function:

```hcl
data "http" "raw" {
  url = "https://raw.githubusercontent.com/nobbs/terraform-provider-sops/refs/heads/main/test/fixtures/raw.sops.txt"
}

output "raw" {
  value = provider::sops::string(data.http.raw.response_body)
}

# Output:
# raw = {
#   "data" = null
#   "raw" = <<-EOT
#   Lorem ipsum dolor sit amet, consectetur adipiscing elit.
#
#   EOT
# }
```
