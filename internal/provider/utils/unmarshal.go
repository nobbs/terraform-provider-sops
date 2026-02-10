// Copyright Alexej Disterhoft <alexej@disterhoft.de> 2024, 2026
// SPDX-License-Identifier: MPL-2.0

package utils

import "fmt"

// UnmarshalDecryptedData unmarshals decrypted data into JSON format based on the specified format.
func UnmarshalDecryptedData(data []byte, format string) (json []byte, err error) {
	switch format {
	case "yaml":
		json, err = ReadYAML(data)
	case "json":
		json, err = ReadJSON(data)
	case "ini":
		json, err = ReadINI(data)
	case "dotenv":
		json, err = ReadENV(data)
	case "binary":
		json, err = []byte{}, nil // we cannot unmarshal binary data, return empty JSON
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal decrypted data: %v", err)
	}

	return json, nil
}
