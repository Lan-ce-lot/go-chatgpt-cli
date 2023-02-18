package cmd

import (
	"bytes"
	"fmt"
	"go-chatgpt-cli/version"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"version"})

	if err := rootCmd.Execute(); err != nil {
		assert.Error(t, err)
		return
	}
	//log.Println(actual.String())
	expected := fmt.Sprintf("go-chatgpt-cli version %s\n", version.Version)
	assert.Equal(t, expected, actual.String(), "actual is not expected")
}
