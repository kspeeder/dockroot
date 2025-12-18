package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogout(t *testing.T) {
	dir := t.TempDir()
	authFile := filepath.Join(dir, "auth.json")
	compatAuthFile := filepath.Join(dir, "config.json")

	// Just a trivial smoke-test exercising one error-handling path.
	// We can’t test full operation without a registry, unit tests should mostly
	// exist in c/common/pkg/auth, not here.
	err := os.WriteFile(authFile, []byte("{}"), 0o700)
	require.NoError(t, err)
	err = os.WriteFile(compatAuthFile, []byte("{}"), 0o700)
	require.NoError(t, err)
	out, err := runSkopeo("logout", "--authfile", authFile, "--compat-auth-file", compatAuthFile, "example.com")
	assertTestFailed(t, out, err, "options for paths to the credential file and to the Docker-compatible credential file can not be set simultaneously")
}
