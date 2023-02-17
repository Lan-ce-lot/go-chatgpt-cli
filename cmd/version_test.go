package cmd

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-chatgpt-cli/version"
	"testing"
)

func TestVersion(t *testing.T) {

	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"version"})
	rootCmd.Execute()
	//log.Println(actual.String())
	expected := fmt.Sprintf("go-chatgpt-cli version %s\n", version.Version)
	assert.Equal(t, expected, actual.String(), "actual is not expected")
}
