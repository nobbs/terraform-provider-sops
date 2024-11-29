terraform {
  required_version = "~> 1.8"

  required_providers {
    sops = {
      source  = "nobbs/sops"
      version = "~> 0.0.1"
    }
  }
}

# There are no configuration options
provider "sops" {}
