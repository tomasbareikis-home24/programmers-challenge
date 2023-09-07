package main

import (
	"strings"
	"testing"

	"github.com/bitfield/script"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	in     string
	expect string
}

func Test(t *testing.T) {
	tests := map[string]testCase{
		"1": {
			in:     "haha",
			expect: "haha",
		},
	}

	for name, tc := range tests {
		t.Run(name, doTest(&tc))
	}
}

func doTest(tc *testCase) func(*testing.T) {
	return func(t *testing.T) {
		s, err := run(tc.in)

		assert.NoError(t, err)

		assert.Equal(
			t,
			strings.TrimSpace(tc.expect),
			strings.TrimSpace(s),
		)
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run("haha")
	}
}

func run(in string) (string, error) {
	return script.Echo(in).Exec("./run.sh").String()
}
