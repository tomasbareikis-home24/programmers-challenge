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
	tests := map[string]testCase{
		"easy": {
			inFile:     "tests/easy/in.txt",
			expectFile: "tests/easy/expected.txt",
		},
		"medium": {
			inFile:     "tests/medium/in.txt",
			expectFile: "tests/medium/expected.txt",
		},
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

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run("haha")
	}
}

func run(inFile string) (string, error) {
	lang := os.Getenv("LANGUAGE")
	return script.File(inFile).Exec("./run-" + lang + ".sh").String()
}
