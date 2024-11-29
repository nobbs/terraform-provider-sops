// Copyright (c) Alexej Disterhoft <alexej@disterhoft.de>
// SPDX-License-Identifier: MPL-2.0

package utils

import "fmt"

// Code returns the given string as an Markdown inline code block.
func Code(s string) string {
	return fmt.Sprintf("`%s`", s)
}
