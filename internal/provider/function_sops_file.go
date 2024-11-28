package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nobbs/terraform-provider-sops/internal/typeutils"
	"github.com/wlevene/ini"
	"gopkg.in/yaml.v3"
)

var sopsFileReturnAttrTypes = map[string]attr.Type{
	"raw":  types.StringType,
	"data": types.DynamicType,
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
			function.StringParameter{
				Name:           "format",
				Description:    "The format of the file to read. If not provided, the format will be inferred from the file extension.",
				AllowNullValue: true,
			},
		},

		Return: function.ObjectReturn{
			AttributeTypes: sopsFileReturnAttrTypes,
		},
	}
}

func (f *SopsFile) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var file, format string

	resp.Error = req.Arguments.Get(ctx, &file, &format)
	if resp.Error != nil {
		return
	}

	// decrypt
	cleartext, err := decrypt.File(file, format)
	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("failed to decrypt file: %v", err))
		return
	}

	var jsonBytes []byte

	switch format {
	case "yaml":
		jsonBytes, err = readYAML(cleartext)
	case "json":
		jsonBytes, err = readJSON(cleartext)
	case "ini":
		jsonBytes, err = readINI(cleartext)
	}

	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("failed to unmarshal decrypted data: %v", err))
		return
	}

	dt, err := typeutils.JSONToDynamicImplied(jsonBytes)
	if err != nil {
		resp.Error = function.NewFuncError(fmt.Sprintf("failed to convert decrypted data to dynamic: %v", err))
		return
	}

	result, diags := types.ObjectValue(
		sopsFileReturnAttrTypes,
		map[string]attr.Value{
			"raw":  types.StringValue(string(cleartext)),
			"data": dt,
		},
	)

	resp.Error = function.FuncErrorFromDiags(ctx, diags)
	if resp.Error != nil {
		return
	}

	resp.Error = resp.Result.Set(ctx, &result)
}

func readYAML(data []byte) ([]byte, error) {
	var v any
	err := yaml.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return json.Marshal(v)
}

func readJSON(data []byte) ([]byte, error) {
	var v any
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return json.Marshal(v)
}

func readINI(data []byte) ([]byte, error) {
	x := ini.New().Load(data)
	if err := x.Err(); err != nil {
		return nil, err
	}

	v := x.Marshal2Json()
	if err := x.Err(); err != nil {
		return nil, err
	}

	return v, nil
}
