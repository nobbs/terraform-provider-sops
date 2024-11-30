output "raw" {
  value = provider::sops::file("./../../../test/fixtures/raw.sops.txt")
}

# raw = {
#   "data" = null
#   "raw" = <<-EOT
#   Lorem ipsum dolor sit amet, consectetur adipiscing elit.
#
#   EOT
# }

output "basic-json" {
  value = provider::sops::file("./../../../test/fixtures/basic.sops.json")
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

output "complex-yaml" {
  value = provider::sops::file("./../../../test/fixtures/complex.sops.yaml")
}

# complex-yaml = {
#   "data" = {
#     "boolean_key" = true
#     "float_key"   = 3.14
#     "integer_key" = 42
#     "list_key" = [
#       "item1",
#       "item2",
#       3,
#       false,
#       null,
#     ]
#     "null_key" = null
#     "object_key" = {
#       "nested_integer" = 100
#       "nested_list" = [
#         "nested item1",
#         200,
#       ]
#       "nested_object" = {
#         "deeper_boolean" = false
#         "deeper_string"  = "deeper example"
#       }
#       "nested_string" = "nested example"
#     }
#     "string_key" = "example string"
#   }
#   "raw" = <<-EOT
#   string_key: example string
#   integer_key: 42
#   float_key: 3.14
#   boolean_key: true
#   null_key: null
#   list_key:
#       - item1
#       - item2
#       - 3
#       - false
#       - null
#   object_key:
#       nested_string: nested example
#       nested_integer: 100
#       nested_list:
#           - nested item1
#           - 200
#       nested_object:
#           deeper_string: deeper example
#           deeper_boolean: false
#
#   EOT
# }
