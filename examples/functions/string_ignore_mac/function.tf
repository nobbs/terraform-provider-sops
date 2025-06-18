data "http" "basic-yaml-mac-mismatch" {
  url = "https://raw.githubusercontent.com/nobbs/terraform-provider-sops/refs/heads/main/test/fixtures/basic-mac-mismatch.sops.yaml"
}

output "basic-yaml-mac-mismatch" {
  value = provider::sops::string_ignore_mac(data.http.basic-yaml-mac-mismatch.response_body, "yaml")
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
