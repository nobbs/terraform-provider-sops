// Copyright (c) Alexej Disterhoft <alexej@disterhoft.de>
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nobbs/terraform-provider-sops/internal/provider/utils"
)

var sopsFileReturnAttrTypes = map[string]attr.Type{
	"raw":  types.StringType,
	"data": types.DynamicType,
}

var _ function.Function = &File{}

type File struct{}

func NewSopsFileFunction() function.Function {
	return &File{}
}

func (f *File) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "file"
}

func (f *File) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Read and decrypt a sops encrypted file into a string",
		Description: "Given a file path to a sops encrypted file, this function will read and decrypt the file into a string.",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "file",
				Description: "The path to the sops encrypted file to read and decrypt.",
			},
		},
		VariadicParameter: function.StringParameter{
			Name:        "format",
			Description: "The format of the encrypted file. Optional, if provided, one of 'yaml', 'json', 'dotenv', 'ini', or 'binary'.",
		},

		Return: function.ObjectReturn{
			AttributeTypes: sopsFileReturnAttrTypes,
		},
	}
}

func (f *File) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
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
	cleartext, err := decrypt.File(file, format)
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
