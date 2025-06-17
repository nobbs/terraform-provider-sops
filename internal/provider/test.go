// Copyright (c) Alexej Disterhoft <alexej@disterhoft.de>
// SPDX-License-Identifier: MPL-2.0

package provider

import "fmt"

const (
	fixture_raw_file                = "test/fixtures/raw.sops.txt"
	fixture_basic_yaml_file         = "test/fixtures/basic.sops.yaml"
	fixture_basic_json_file         = "test/fixtures/basic.sops.json"
	fixture_complex_yaml_file       = "test/fixtures/complex.sops.yaml"
	fixture_complex_json_file       = "test/fixtures/complex.sops.json"
	fixture_sample_ini_file         = "test/fixtures/sample.sops.ini"
	fixture_sample_env_file         = "test/fixtures/dot.sops.env"
	fixture_basic_mac_mismatch_file = "test/fixtures/basic-mac-mismatch.sops.yaml"
)

var (
	functions = map[string]string{
		"file":              "provider::sops::file",
		"file_ignore_mac":   "provider::sops::file_ignore_mac",
		"string":            "provider::sops::string",
		"string_ignore_mac": "provider::sops::string_ignore_mac",
	}
)

func testHelperFunctionConfig(fn string, file string, format string) string {
	if format != "" {
		format = fmt.Sprintf(", %q", format)
	}

	switch fn {
	case "file", "file_ignore_mac":
		return fmt.Sprintf(
			`
output "test" {
	value = %s("%s"%s)
}
`,
			functions[fn], file, format,
		)
	case "string", "string_ignore_mac":
		return fmt.Sprintf(
			`
output "test" {
	value = %s(file("%s")%s)
}
`,
			functions[fn], file, format,
		)
	}

	return ""
}
