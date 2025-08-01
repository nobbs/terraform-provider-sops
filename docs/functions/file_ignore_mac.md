---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "file_ignore_mac function - sops"
subcategory: ""
description: |-
  Reads and decrypts a sops encrypted file ignoring MAC mismatch.
---

# function: file_ignore_mac

Reads and decrypts a [sops](https://getsops.io/) encrypted file ignoring MAC mismatch. An optional format can be
provided to specify the format of the encrypted file. If not provided, we will try to infer
the format from the file extension. Supported formats are `yaml`, `json`, `dotenv`, `ini`, and `binary`.

If the file format is any of the supported formats other than `binary`, the
decrypted data will also be returned as an object in the `data` attribute.
Regardless of the format, the raw decrypted data will always be returned in the `raw` attribute.

Decryption is based on the sops library, so it will use the same heuristics and key sources
as sops to attempt to decrypt the file.

## Example Usage

```terraform
output "basic-yaml-mac-mismatch" {
  value = provider::sops::file_ignore_mac("./../../../test/fixtures/basic-mac-mismatch.sops.yaml")
}

# basic-yaml-mac-mismatch = {
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

## Signature

<!-- signature generated by tfplugindocs -->
```text
file_ignore_mac(file string, format string...) object
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `file` (String) The path to the sops encrypted file.
<!-- variadic argument generated by tfplugindocs -->
1. `format` (Variadic, String, Nullable) The format of the encrypted file. Supported formats are `yaml`, `json`, `dotenv`, `ini`, and `binary`. Optional.
