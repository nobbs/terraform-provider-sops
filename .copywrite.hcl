schema_version = 1

project {
  license        = "MPL-2.0"
  copyright_year = 2024
  copyright_holder = "Alexej Disterhoft <alexej@disterhoft.de>"

  header_ignore = [
    # examples used within documentation (prose)
    "examples/**",

    # GitHub issue template configuration
    ".github/ISSUE_TEMPLATE/*.yml",

    # golangci-lint tooling configuration
    ".golangci.yml",

    # GoReleaser tooling configuration
    ".goreleaser.yml",

    # encrypted test fixtures
    "test/fixtures/**",
  ]
}
