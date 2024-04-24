package files

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	basePath := "basePath"

	testStorage := Storage{basePath: basePath}

	require.Equal(t, testStorage, New(basePath))
}
