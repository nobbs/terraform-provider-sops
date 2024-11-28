// Copyright (c) Microsoft Corporation
// SPDX-License-Identifier: MPL-2.0

package utils

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const nullStr = "null"

// JSONToDynamicImplied is similar to FromJSON, while it is for typeless case.
// In which case, the following type conversion rules are applied (Go -> TF):
// - bool: bool
// - float64: number
// - string: string
// - []interface{}: tuple
// - map[string]interface{}: object
// - nil: null (dynamic)
// In case the input json is of zero-length, it returns null (dynamic).
func JSONToDynamicImplied(b []byte) (types.Dynamic, error) {
	if len(b) == 0 {
		return types.DynamicNull(), nil
	}

	_, v, err := attrValueFromJSONImplied(b)
	if err != nil {
		return types.Dynamic{}, err
	}

	return types.DynamicValue(v), nil
}

func attrValueFromJSONImplied(b []byte) (attr.Type, attr.Value, error) {
	if string(b) == nullStr {
		return types.DynamicType, types.DynamicNull(), nil
	}

	var object map[string]json.RawMessage
	if err := json.Unmarshal(b, &object); err == nil {
		attrTypes := map[string]attr.Type{}
		attrVals := map[string]attr.Value{}

		for k, v := range object {
			attrTypes[k], attrVals[k], err = attrValueFromJSONImplied(v)
			if err != nil {
				return nil, nil, err
			}
		}

		typ := types.ObjectType{AttrTypes: attrTypes}
		val, diags := types.ObjectValue(attrTypes, attrVals)

		if diags.HasError() {
			diag := diags.Errors()[0]

			return nil, nil, fmt.Errorf("%s: %s", diag.Summary(), diag.Detail())
		}

		return typ, val, nil
	}

	var array []json.RawMessage
	if err := json.Unmarshal(b, &array); err == nil {
		eTypes := []attr.Type{}
		eVals := []attr.Value{}

		for _, e := range array {
			eType, eVal, err := attrValueFromJSONImplied(e)
			if err != nil {
				return nil, nil, err
			}

			eTypes = append(eTypes, eType)
			eVals = append(eVals, eVal)
		}

		typ := types.TupleType{ElemTypes: eTypes}
		val, diags := types.TupleValue(eTypes, eVals)

		if diags.HasError() {
			diag := diags.Errors()[0]

			return nil, nil, fmt.Errorf("%s: %s", diag.Summary(), diag.Detail())
		}

		return typ, val, nil
	}

	// Primitives
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal %s: %w", string(b), err)
	}

	switch v := v.(type) {
	case bool:
		return types.BoolType, types.BoolValue(v), nil
	case float64:
		return types.NumberType, types.NumberValue(big.NewFloat(v)), nil
	case string:
		return types.StringType, types.StringValue(v), nil
	case nil:
		return types.DynamicType, types.DynamicNull(), nil
	default:
		return nil, nil, fmt.Errorf("unhandled type: %T", v)
	}
}
