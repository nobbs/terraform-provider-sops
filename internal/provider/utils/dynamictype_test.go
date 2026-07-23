// Copyright Alexej Disterhoft <alexej@disterhoft.de> 2024, 2026
// SPDX-License-Identifier: MPL-2.0

package utils

import (
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestJSONToDynamicImpliedPreservesNumberPrecision(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		number string
	}{
		{
			name:   "integer beyond float64 exact range",
			input:  `{"number":9007199254740993}`,
			number: "9007199254740993",
		},
		{
			name:   "decimal beyond float64 precision",
			input:  `{"number":0.12345678901234567890123456789}`,
			number: "0.12345678901234567890123456789",
		},
		{
			name:   "number beyond float64 range",
			input:  `{"number":1e1000}`,
			number: "1e1000",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			jsonData, err := ReadJSON([]byte(test.input))
			if err != nil {
				t.Fatalf("ReadJSON() error = %v", err)
			}

			dynamicValue, err := JSONToDynamicImplied(jsonData)
			if err != nil {
				t.Fatalf("JSONToDynamicImplied() error = %v", err)
			}

			objectValue, ok := dynamicValue.UnderlyingValue().(types.Object)
			if !ok {
				t.Fatalf("JSONToDynamicImplied() underlying value type = %T, want types.Object", dynamicValue.UnderlyingValue())
			}

			numberValue, ok := objectValue.Attributes()["number"].(types.Number)
			if !ok {
				t.Fatalf("number attribute type = %T, want types.Number", objectValue.Attributes()["number"])
			}

			expected, _, err := big.ParseFloat(test.number, 10, 512, big.ToNearestEven)
			if err != nil {
				t.Fatalf("big.ParseFloat() error = %v", err)
			}

			if numberValue.ValueBigFloat().Cmp(expected) != 0 {
				t.Errorf(
					"number value = %s, want %s",
					numberValue.ValueBigFloat().Text('g', -1),
					expected.Text('g', -1),
				)
			}
		})
	}
}
