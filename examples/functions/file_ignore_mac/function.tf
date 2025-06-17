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
