package configs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckerE2E(t *testing.T) {
	t.Skip()
	checker := defaultBitbucketClient{}
	require.NoError(t, checker.TestConnection("https://mybitbucket.com/rest", "XXX"))
}
