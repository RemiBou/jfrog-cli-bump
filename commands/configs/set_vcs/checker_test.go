package set_vcs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckerE2E(t *testing.T) {
	t.Skip()
	checker := defaultVcsConfigChecker{}
	require.NoError(t, checker.check(vcsConfig{
		Url:   "https://mybitbucket.com/rest",
		Token: "XXX",
	}))
}
