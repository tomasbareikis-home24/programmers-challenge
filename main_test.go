package main

import (
	"os"
	"strings"
	"testing"

	"github.com/bitfield/script"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	inFile     string
	expectFile string
}

func Test(t *testing.T) {
	tests := map[string]testCase{}

	dirs := []string{"easy", "empty", "one", "noone", "meter", "tiny"}
	for _, d := range dirs {
		tests[d] = testCase{
			inFile:     "tests/" + d + "/in.txt",
			expectFile: "tests/" + d + "/expected.txt",
		}
	}

	for name, tc := range tests {
		t.Run(name, doTest(&tc))
	}
}

func doTest(tc *testCase) func(*testing.T) {
	return func(t *testing.T) {
		expectedOutput, err := script.File(tc.expectFile).String()
		assert.NoError(t, err)

		output, err := run(tc.inFile)
		assert.NoError(t, err)

		assert.Equal(
			t,
			strings.TrimSpace(expectedOutput),
			strings.TrimSpace(output),
		)
	}
}

func run(inFile string) (string, error) {
	lang := os.Getenv("LANGUAGE")
	return script.File(inFile).Exec("./run-" + lang + ".sh").String()
}
