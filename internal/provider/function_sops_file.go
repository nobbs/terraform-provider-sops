package provider

import (
	"context"
	"fmt"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var sopsFileReturnAttrTypes = map[string]attr.Type{
	"raw": types.StringType,
}

var _ function.Function = &SopsFile{}

type SopsFile struct{}

func NewSopsFileFunction() function.Function {
	return &SopsFile{}
}

func (f *SopsFile) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "sops_file"
}

func (f *SopsFile) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Read and decrypt a sops encrypted file into a string",
		Description: "Given a file path to a sops encrypted file, this function will read and decrypt the file into a string.",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "file",
				Description: "The path to the sops encrypted file to read and decrypt.",
			},
		},

		Return: function.ObjectReturn{
			AttributeTypes: sopsFileReturnAttrTypes,
		},
	}
}

func (f *SopsFile) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var file string

	resp.Error = req.Arguments.Get(ctx, &file)
	if resp.Error != nil {
		return
	}

	// decrypt
	cleartext, err := decrypt.File(file, "binary")
	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("failed to decrypt file: %v", err))
		return
	}

	result, diags := types.ObjectValue(
		sopsFileReturnAttrTypes,
		map[string]attr.Value{
			"raw": types.StringValue(string(cleartext)),
		},
	)

	resp.Error = function.FuncErrorFromDiags(ctx, diags)
	if resp.Error != nil {
		return
	}

	resp.Error = resp.Result.Set(ctx, &result)
}
