// Copyright (c) Alexej Disterhoft <alexej@disterhoft.de>
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/lithammer/dedent"
	"github.com/nobbs/terraform-provider-sops/internal/provider/utils"
)

var sopsStringReturnAttrTypes = map[string]attr.Type{
	"raw":  types.StringType,
	"data": types.DynamicType,
}

var _ function.Function = &stringFunction{}

type stringFunction struct{}

func NewStringFunction() function.Function {
	return &stringFunction{}
}

func (f *stringFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "string"
}

func (f *stringFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Reads and decrypts a sops encrypted string.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
			Reads and decrypts a [sops](https://getsops.io/) encrypted string.
			An optional format can be provided to specify the format of the encrypted string. If not
			provided, structured data will not be automatically converted to an object. Supported formats
			are ` + utils.Code("yaml") + `, ` + utils.Code("json") + `, ` + utils.Code("dotenv") + `, ` +
			utils.Code("ini") + `, and ` + utils.Code("binary") + `.

			If the data format is any of the supported formats other than ` + utils.Code("binary") + `, the
			decrypted data will also be returned as an object in the ` + utils.Code("data") + ` attribute.
			Regardless of the format, the raw decrypted data will always be returned in the ` +
			utils.Code("raw") + ` attribute.

			Decryption is based on the sops library, so it will use the same heuristics and key sources
			as sops to attempt to decrypt the data.
			`)),

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "data",
				MarkdownDescription: "The sops encrypted string.",
			},
		},
		VariadicParameter: function.StringParameter{
			Name:           "format",
			Description:    "The format of the encrypted file. Supported formats are `yaml`, `json`, `dotenv`, `ini`, and `binary`. Optional.",
			AllowNullValue: true,
		},

		Return: function.ObjectReturn{
			AttributeTypes: sopsStringReturnAttrTypes,
		},
	}
}

func (f *stringFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data string
	var varargs []string

	resp.Error = req.Arguments.Get(ctx, &data, &varargs)
	if resp.Error != nil {
		return
	}

	// infer format from file extension if not explicitly provided
	format := ""
	if len(varargs) > 0 {
		format = varargs[0]
	} else {
		format = "binary"
	}

	if !utils.IsValidFormat(format) {
		resp.Error = function.NewFuncError(fmt.Sprintf("invalid format: %s", format))
		return
	}

	// decrypt sops file
	databytes := []byte(data)
	cleartext, err := decrypt.Data(databytes, format)
	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("failed to decrypt file: %v", err))
		return
	}

	var json []byte

	switch format {
	case "yaml":
		json, err = utils.ReadYAML(cleartext)
	case "json":
		json, err = utils.ReadJSON(cleartext)
	case "ini":
		json, err = utils.ReadINI(cleartext)
	case "dotenv":
		json, err = utils.ReadENV(cleartext)
	}

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
		sopsFileReturnAttrTypes,
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
