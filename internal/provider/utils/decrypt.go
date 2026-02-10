// Copyright Alexej Disterhoft <alexej@disterhoft.de> 2024, 2026
// SPDX-License-Identifier: MPL-2.0

package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/getsops/sops/v3/aes"
	"github.com/getsops/sops/v3/cmd/sops/common"
	"github.com/getsops/sops/v3/cmd/sops/formats"
	"github.com/getsops/sops/v3/config"
)

// DecryptOptions contains options for the Decrypt function.
type DecryptOptions struct {
	// IgnoreMACMismatch indicates whether to ignore MAC mismatch errors.
	IgnoreMACMismatch bool
}

// decrypt decrypts the given data using the specified format.
//
// This function is mostly taken from the sops codebase and modified to allow ignoring MAC mismatch
// errors, henceforth the function is following the license of the sops codebase, i.e.
// MPL-2.0.
func decrypt(data []byte, format formats.Format, opts DecryptOptions) (cleartext []byte, err error) {
	store := common.StoreForFormat(format, config.NewStoresConfig())

	// Load SOPS file and access the data key
	tree, err := store.LoadEncryptedFile(data)
	if err != nil {
		return nil, err
	}
	key, err := tree.Metadata.GetDataKey()
	if err != nil {
		return nil, err
	}

	// Decrypt the tree
	cipher := aes.NewCipher()
	mac, err := tree.Decrypt(key, cipher)
	if err != nil {
		return nil, err
	}

	// Compute the original MAC of the tree if not ignoring MAC mismatch
	if !opts.IgnoreMACMismatch {
		// Compute the hash of the cleartext tree and compare it with
		// the one that was stored in the document. If they match,
		// integrity was preserved
		originalMac, err := cipher.Decrypt(
			tree.Metadata.MessageAuthenticationCode,
			key,
			tree.Metadata.LastModified.Format(time.RFC3339),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt original mac: %w", err)
		}
		if originalMac != mac {
			return nil, fmt.Errorf("failed to verify data integrity. expected mac %q, got %q", originalMac, mac)
		}
	}

	return store.EmitPlainFile(tree.Branches)
}

// DecryptData decrypts the given data using the specified format and options.
func DecryptData(data []byte, format string, opts DecryptOptions) (cleartext []byte, err error) {
	formatEnum := formats.FormatFromString(format)
	return decrypt(data, formatEnum, opts)
}

// DecryptFile decrypts the file at the given path using the specified format and options.
func DecryptFile(path string, format string, opts DecryptOptions) (cleartext []byte, err error) {
	// Read the file into an []byte
	encryptedData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read %q: %w", path, err)
	}

	formatEnum := formats.FormatFromString(format)
	return decrypt(encryptedData, formatEnum, opts)
}
