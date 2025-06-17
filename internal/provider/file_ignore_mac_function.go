// Copyright (c) Alexej Disterhoft <alexej@disterhoft.de>
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/lithammer/dedent"
	"github.com/nobbs/terraform-provider-sops/internal/provider/utils"
)

var sopsFileIgnoreMacReturnAttrTypes = map[string]attr.Type{
	"raw":  types.StringType,
	"data": types.DynamicType,
}

// Ensure that fileIgnoreMacFunction implements the Function interface.
var _ function.Function = &fileIgnoreMacFunction{}

type fileIgnoreMacFunction struct{}

func NewFileIgnoreMacFunction() function.Function {
	return &fileIgnoreMacFunction{}
}

func (f *fileIgnoreMacFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "file_ignore_mac"
}

func (f *fileIgnoreMacFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Reads and decrypts a sops encrypted file ignoring MAC mismatch.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
			Reads and decrypts a [sops](https://getsops.io/) encrypted file ignoring MAC mismatch. An optional format can be
			provided to specify the format of the encrypted file. If not provided, we will try to infer
			the format from the file extension. Supported formats are ` + utils.Code("yaml") + `, ` +
			utils.Code("json") + `, ` + utils.Code("dotenv") + `, ` + utils.Code("ini") + `, and ` +
			utils.Code("binary") + `.

			If the file format is any of the supported formats other than ` + utils.Code("binary") + `, the
			decrypted data will also be returned as an object in the ` + utils.Code("data") + ` attribute.
			Regardless of the format, the raw decrypted data will always be returned in the ` +
			utils.Code("raw") + ` attribute.

			Decryption is based on the sops library, so it will use the same heuristics and key sources
			as sops to attempt to decrypt the file.
		`)),

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "file",
				MarkdownDescription: "The path to the sops encrypted file.",
			},
		},
		VariadicParameter: function.StringParameter{
			Name:           "format",
			Description:    "The format of the encrypted file. Supported formats are `yaml`, `json`, `dotenv`, `ini`, and `binary`. Optional.",
			AllowNullValue: true,
		},

		Return: function.ObjectReturn{
			AttributeTypes: sopsFileIgnoreMacReturnAttrTypes,
		},
	}
}

func (f *fileIgnoreMacFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var file string
	var varargs []string

	resp.Error = req.Arguments.Get(ctx, &file, &varargs)
	if resp.Error != nil {
		return
	}

	// infer format from file extension if not explicitly provided
	format := ""
	if len(varargs) > 0 {
		format = varargs[0]
	} else {
		format = utils.FileFormatFromPath(file)
	}

	if !utils.IsValidFormat(format) {
		resp.Error = function.NewFuncError(fmt.Sprintf("invalid format: %s", format))
		return
	}

	// decrypt sops file
	cleartext, err := utils.DecryptFile(file, format, utils.DecryptOptions{IgnoreMACMismatch: true})
	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("failed to decrypt file: %v", err))
		return
	}

	json, err := utils.UnmarshalDecryptedData(cleartext, format)
	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("failed to unmarshal decrypted data: %v", err))
		return
	}

	dynamicData, err := utils.JSONToDynamicImplied(json)
	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("failed to convert decrypted data to dynamic data: %v", err))
		return
	}

	result, diags := types.ObjectValue(
		sopsFileIgnoreMacReturnAttrTypes,
		map[string]attr.Value{
			"raw":  types.StringValue(string(cleartext)),
			"data": dynamicData,
		},
	)

	resp.Error = function.FuncErrorFromDiags(ctx, diags)
	if resp.Error != nil {
		return
	}

	resp.Error = resp.Result.Set(ctx, &result)
}
