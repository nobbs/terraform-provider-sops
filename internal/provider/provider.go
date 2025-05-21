// Copyright (c) Alexej Disterhoft <alexej@disterhoft.de>
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/lithammer/dedent"
	"github.com/nobbs/terraform-provider-sops/internal/provider/utils"
)

// Ensure SopsProvider satisfies various provider interfaces.
var _ provider.Provider = &SopsProvider{}
var _ provider.ProviderWithFunctions = &SopsProvider{}

// SopsProvider defines the provider implementation.
type SopsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *SopsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "sops"
	resp.Version = p.version
}

func (p *SopsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
			The Sops provider offers functions to work with [sops](https://getsops.io/) encrypted files.

			This provider is inspired by the
			[carlpett/terraform-provider-sops](https://registry.terraform.io/providers/carlpett/sops)
			provider, but instead of using data sources, it implements its functionality as functions.
			This approach provides more flexibility and ensures that decrypted data is not stored in the
			state file.

			Moreover, if the decrypted data is in one of the supported formats (` + utils.Code("yaml") + `, ` +
			utils.Code("json") + `, ` + utils.Code("dotenv") + `, ` + utils.Code("ini") + `), it will also be
			returned as a nested object in the ` + utils.Code("data") + ` attribute. This allows for easier
			access to specific values within structured data.
		`)),
	}
}

func (p *SopsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// no configuration needed
}

func (p *SopsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *SopsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *SopsProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewFileFunction,
		NewStringFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SopsProvider{
			version: version,
		}
	}
}
