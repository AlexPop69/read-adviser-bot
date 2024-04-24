package telegram

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	host := "host.host"
	token := "testToken"

	testClient := &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}

	require.Equal(t, New(host, token), testClient)
}

func Test_newBasePath(t *testing.T) {
	//Arrange
	token := "testToken"
	expected := "bot" + token

	//Act
	result := newBasePath(token)

	// Assert
	if result != expected {
		t.Errorf("incorresct result. Have: %s, Want: %s", result, expected)
	}

	assert.Equal(t, expected, newBasePath(token), fmt.Sprintf("Have: %s, Want: %s", result, expected))
}
