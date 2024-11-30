// Copyright (c) Alexej Disterhoft <alexej@disterhoft.de>
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

const (
	stringFunctionConfig_raw_file          = "test/fixtures/raw.sops.txt"
	stringFunctionConfig_basic_yaml_file   = "test/fixtures/basic.sops.yaml"
	stringFunctionConfig_basic_json_file   = "test/fixtures/basic.sops.json"
	stringFunctionConfig_complex_yaml_file = "test/fixtures/complex.sops.yaml"
	stringFunctionConfig_complex_json_file = "test/fixtures/complex.sops.json"
	stringFunctionConfig_sample_ini_file   = "test/fixtures/sample.sops.ini"
	stringFunctionConfig_sample_env_file   = "test/fixtures/dot.sops.env"
)

func helperStringFunctionConfig(file string, format string) string {
	if format != "" {
		format = fmt.Sprintf(", %q", format)
	}

	return fmt.Sprintf(
		`
output "test" {
	value = provider::sops::string(file("%s")%s)
}
`,
		file, format,
	)
}

func TestStringFunction_raw(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	fixture := fmt.Sprintf("%s/../../%s", wd, stringFunctionConfig_raw_file)

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: helperStringFunctionConfig(fixture, "binary"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"raw": knownvalue.StringExact("Lorem ipsum dolor sit amet, consectetur adipiscing elit.\n"),
						}),
					),
				},
			},
			{
				Config: helperStringFunctionConfig(fixture, "binary"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"raw": knownvalue.StringExact("Lorem ipsum dolor sit amet, consectetur adipiscing elit.\n"),
						}),
					),
				},
			},
		},
	})
}

func TestStringFunction_basic_yaml(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	fixture := fmt.Sprintf("%s/../../%s", wd, stringFunctionConfig_basic_yaml_file)

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: helperStringFunctionConfig(fixture, "yaml"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"data": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"abc":      knownvalue.StringExact("xyz"),
								"integers": knownvalue.Int64Exact(123),
								"truthy":   knownvalue.Bool(true),
								"floats":   knownvalue.Float64Exact(3.14e-10),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestStringFunction_basic_json(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	fixture := fmt.Sprintf("%s/../../%s", wd, stringFunctionConfig_basic_json_file)

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: helperStringFunctionConfig(fixture, "json"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"data": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"abc":      knownvalue.StringExact("xyz"),
								"integers": knownvalue.Int64Exact(123),
								"truthy":   knownvalue.Bool(true),
								"floats":   knownvalue.Float64Exact(3.14e-10),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestStringFunction_complex_yaml(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	fixture := fmt.Sprintf("%s/../../%s", wd, stringFunctionConfig_complex_yaml_file)

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: helperStringFunctionConfig(fixture, "yaml"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"data": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"string_key":  knownvalue.StringExact("example string"),
								"integer_key": knownvalue.Int64Exact(42),
								"float_key":   knownvalue.Float64Exact(3.14),
								"boolean_key": knownvalue.Bool(true),
								"null_key":    knownvalue.Null(),
								"list_key": knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact("item1"),
									knownvalue.StringExact("item2"),
									knownvalue.Int64Exact(3),
									knownvalue.Bool(false),
									knownvalue.Null(),
								}),
								"object_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"nested_string":  knownvalue.StringExact("nested example"),
									"nested_integer": knownvalue.Int64Exact(100),
									"nested_list": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("nested item1"),
										knownvalue.Int64Exact(200),
									}),
									"nested_object": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"deeper_string":  knownvalue.StringExact("deeper example"),
										"deeper_boolean": knownvalue.Bool(false),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestStringFunction_complex_json(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	fixture := fmt.Sprintf("%s/../../%s", wd, stringFunctionConfig_complex_json_file)

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: helperStringFunctionConfig(fixture, "json"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"data": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"string_key":  knownvalue.StringExact("example string"),
								"integer_key": knownvalue.Int64Exact(42),
								"float_key":   knownvalue.Float64Exact(3.14),
								"boolean_key": knownvalue.Bool(true),
								"null_key":    knownvalue.Null(),
								"list_key": knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact("item1"),
									knownvalue.StringExact("item2"),
									knownvalue.Int64Exact(3),
									knownvalue.Bool(false),
									knownvalue.Null(),
								}),
								"object_key": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"nested_string":  knownvalue.StringExact("nested example"),
									"nested_integer": knownvalue.Int64Exact(100),
									"nested_list": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.StringExact("nested item1"),
										knownvalue.Int64Exact(200),
									}),
									"nested_object": knownvalue.ObjectExact(map[string]knownvalue.Check{
										"deeper_string":  knownvalue.StringExact("deeper example"),
										"deeper_boolean": knownvalue.Bool(false),
									}),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestStringFunction_sample_ini(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	fixture := fmt.Sprintf("%s/../../%s", wd, stringFunctionConfig_sample_ini_file)

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: helperStringFunctionConfig(fixture, "ini"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"data": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"general": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"appname": knownvalue.StringExact("SampleApp"),
									"version": knownvalue.StringExact("1.0"),
								}),
								"database": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"host":     knownvalue.StringExact("localhost"),
									"port":     knownvalue.StringExact("3306"),
									"username": knownvalue.StringExact("root"),
									"password": knownvalue.StringExact("password"),
								}),
								"logging": knownvalue.ObjectExact(map[string]knownvalue.Check{
									"level": knownvalue.StringExact("DEBUG"),
									"file":  knownvalue.StringExact("/var/log/sampleapp.log"),
								}),
							}),
						}),
					),
				},
			},
		},
	})
}

func TestStringFunction_sample_env(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	fixture := fmt.Sprintf("%s/../../%s", wd, stringFunctionConfig_sample_env_file)

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(version.Must(version.NewVersion("1.8.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: helperStringFunctionConfig(fixture, "dotenv"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.ObjectPartial(map[string]knownvalue.Check{
							"data": knownvalue.ObjectExact(map[string]knownvalue.Check{
								"HELLO":        knownvalue.StringExact("world"),
								"SOME_NUMBER":  knownvalue.StringExact("42"),
								"DATABASE_URL": knownvalue.StringExact("postgres://localhost:5432/mydb"),
							}),
						}),
					),
				},
			},
		},
	})
}
