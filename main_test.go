package main

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMustToken(t *testing.T) {
	OkToken := flag.String(
		"tg-bot-token",
		"testtoken",
		"token for access to telegram bot",
	)

	flag.Parse()

	assert.NotEqual(t, "", *OkToken)
	assert.Equal(t, "testtoken", *OkToken)

}

func Test_mustHost(t *testing.T) {
	expected := "api.telegram.org"

	result := mustHost()

	require.Equal(t, expected, result, fmt.Sprintf(`want: %s
have: %s`, expected, result))
}
