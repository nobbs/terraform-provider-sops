data "http" "raw" {
  url = "https://raw.githubusercontent.com/nobbs/terraform-provider-sops/refs/heads/main/test/fixtures/raw.sops.txt"
}

output "raw" {
  value = provider::sops::string(data.http.raw.response_body)
}

# raw = {
#   "data" = null
#   "raw" = <<-EOT
#   Lorem ipsum dolor sit amet, consectetur adipiscing elit.
#
#   EOT
# }

output "basic-json" {
  value = provider::sops::string(file("./../../../test/fixtures/basic.sops.json"), "json")
}

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
