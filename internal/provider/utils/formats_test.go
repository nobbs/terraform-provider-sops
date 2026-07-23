// Copyright Alexej Disterhoft <alexej@disterhoft.de> 2024, 2026
// SPDX-License-Identifier: MPL-2.0

package utils

import (
	"encoding/json"
	"testing"
)

func TestDecodeJSONPreservingNumbersAllowsTrailingWhitespace(t *testing.T) {
	t.Parallel()

	value, err := decodeJSONPreservingNumbers([]byte("{\"number\":9007199254740993}\n\t "))
	if err != nil {
		t.Fatalf("decodeJSONPreservingNumbers() error = %v", err)
	}

	object, ok := value.(map[string]any)
	if !ok {
		t.Fatalf("decodeJSONPreservingNumbers() type = %T, want map[string]any", value)
	}

	number, ok := object["number"].(json.Number)
	if !ok {
		t.Fatalf("number type = %T, want json.Number", object["number"])
	}

	if got, want := number.String(), "9007199254740993"; got != want {
		t.Errorf("number = %q, want %q", got, want)
	}
}

func TestDecodeJSONPreservingNumbersRejectsTrailingData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "second object",
			input: `{} {}`,
		},
		{
			name:  "invalid token",
			input: `{} invalid`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, gotErr := decodeJSONPreservingNumbers([]byte(test.input))

			var value any
			wantErr := json.Unmarshal([]byte(test.input), &value)
			if gotErr == nil {
				t.Fatal("decodeJSONPreservingNumbers() error = nil, want error")
			}
			if gotErr.Error() != wantErr.Error() {
				t.Errorf(
					"decodeJSONPreservingNumbers() error = %q, want %q",
					gotErr,
					wantErr,
				)
			}
		})
	}
}

func BenchmarkDecodeJSONPreservingNumbers(b *testing.B) {
	data := []byte(
		`{"number":9007199254740993,"decimal":0.12345678901234567890123456789,"values":[1,2,3,4,5]}`,
	)

	for b.Loop() {
		if _, err := decodeJSONPreservingNumbers(data); err != nil {
			b.Fatal(err)
		}
	}
}
