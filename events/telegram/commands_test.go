package telegram

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isURL(t *testing.T) {
	okURL := "https://github.com/AlexPop69/"
	u, err := url.Parse(okURL)

	require.NotEqual(t, "", u.Host)
	require.Equal(t, "github.com", u.Host)

	require.Equal(t, nil, err)

	require.Equal(t, true, isURL(okURL))

	wrongURL := "htfgs:/github,com/AlexPop69/"
	require.Equal(t, false, isURL(wrongURL))
}

func Test_isAddCmd(t *testing.T) {
	text := "https://github.com/AlexPop69/"

	require.Equal(t, true, isURL(text))
}
