//go:build piv

/*
Copyright 2022 Gravitational, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package keys_test

import (
	"context"
	"crypto/rand"
	"crypto/x509/pkix"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gravitational/teleport/api/utils/keys"
	"github.com/gravitational/teleport/api/utils/prompt"
)

// TestGetOrGenerateYubiKeyPrivateKey tests GetOrGenerateYubiKeyPrivateKey.
func TestGetOrGenerateYubiKeyPrivateKey(t *testing.T) {
	// This test expects a yubiKey to be connected with default PIV
	// settings and will overwrite any PIV data on the yubiKey.
	if os.Getenv("TELEPORT_TEST_YUBIKEY_PIV") == "" {
		t.Skipf("Skipping TestGenerateYubiKeyPrivateKey because TELEPORT_TEST_YUBIKEY_PIV is not set")
	}

	ctx := context.Background()
	resetYubikey(ctx, t)

	// Generate a new YubiKeyPrivateKey.
	priv, err := keys.GetOrGenerateYubiKeyPrivateKey(false)
	require.NoError(t, err)

	// Test creating a self signed certificate with the key.
	digest := make([]byte, 32)
	_, err = priv.Sign(rand.Reader, digest, nil)
	require.NoError(t, err)

	// Another call to GetOrGenerateYubiKeyPrivateKey should retrieve the previously generated key.
	retrievePriv, err := keys.GetOrGenerateYubiKeyPrivateKey(false)
	require.NoError(t, err)
	require.Equal(t, priv.Public(), retrievePriv.Public())

	// parsing the key's private key PEM should produce the same key as well.
	retrievePriv, err = keys.ParsePrivateKey(priv.PrivateKeyPEM())
	require.NoError(t, err)
	require.Equal(t, priv.Public(), retrievePriv.Public())
}

func TestOverwritePrompt(t *testing.T) {
	// This test expects a yubiKey to be connected with default PIV
	// settings and will overwrite any PIV data on the yubiKey.
	if os.Getenv("TELEPORT_TEST_YUBIKEY_PIV") == "" {
		t.Skipf("Skipping TestGenerateYubiKeyPrivateKey because TELEPORT_TEST_YUBIKEY_PIV is not set")
	}

	ctx := context.Background()
	resetYubikey(ctx, t)

	oldStdin := prompt.Stdin()
	t.Cleanup(func() { prompt.SetStdin(oldStdin) })

	slot, err := keys.GetDefaultKeySlot(keys.PrivateKeyPolicyHardwareKeyTouch)
	require.NoError(t, err)

	testOverwritePrompt := func(t *testing.T) {
		// Fail to overwrite slot when user denies
		prompt.SetStdin(prompt.NewFakeReader().AddString("n"))
		_, err := keys.GetOrGenerateYubiKeyPrivateKey(true)
		require.Error(t, err)

		// Successfully overwrite slot when user accepts
		prompt.SetStdin(prompt.NewFakeReader().AddString("y"))
		_, err = keys.GetOrGenerateYubiKeyPrivateKey(true)
		require.NoError(t, err)
	}

	t.Run("invalid metadata cert", func(t *testing.T) {
		t.Cleanup(func() { resetYubikey(ctx, t) })

		// Set a non-teleport certificate in the slot.
		y, err := keys.FindYubiKey(0)
		require.NoError(t, err)
		err = y.SetMetadataCertificate(slot, pkix.Name{Organization: []string{"not-teleport"}})
		require.NoError(t, err)

		testOverwritePrompt(t)
	})

	t.Run("invalid key policies", func(t *testing.T) {
		t.Cleanup(func() { resetYubikey(ctx, t) })

		// Generate a key that does not require touch in the slot that Teleport expects to require touch.
		y, err := keys.FindYubiKey(0)
		require.NoError(t, err)
		_, err = y.GeneratePrivateKey(slot, keys.PrivateKeyPolicyHardwareKey)
		require.NoError(t, err)

		testOverwritePrompt(t)
	})
}

// resetYubikey connects to the first yubiKey and resets it to defaults.
func resetYubikey(ctx context.Context, t *testing.T) {
	t.Helper()
	y, err := keys.FindYubiKey(0)
	require.NoError(t, err)
	require.NoError(t, y.Reset())
}
