package rules_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonda2208/go-rules"
	"github.com/stretchr/testify/assert"
)

func TestAntlrParser(t *testing.T) {
	testcases := []struct {
		rules   string
		isError bool
	}{
		{
			`a == false`,
			false,
		},
		{
			`a == 0`,
			false,
		},
		{
			`a == -1`,
			false,
		},
		{
			`a == +1`,
			false,
		},
		{
			`a == "0"`,
			false,
		},
		{
			`a == 2h`,
			false,
		},
		{
			`a == strlen("bc")`,
			false,
		},
		{
			`a == "2006-01-02T15:04:05Z"`,
			false,
		},
		{
			`a in [0, 1, 2]`,
			false,
		},
		{
			`a in ["0", "1", "2"]`,
			false,
		},
		{
			`a == 0 && (b == 1 || c == 2)`,
			false,
		},
		{
			`(a == 0 && b == 1) || (c == 2 && d == 3)`,
			false,
		},
		{
			`a = 1`,
			true,
		},
		{
			`a == 1hh`,
			true,
		},
		{
			`a == 1u`,
			true,
		},
		{
			`(a == 1 && b == 2) || (c == 3 && d == 4`,
			true,
		},
		{
			`a in []`,
			true,
		},
		{
			`a in [0, "1", 2]`,
			true,
		},
		{
			`a in [true, false]`,
			true,
		},
		{
			`a in [0, 1, false]`,
			true,
		},
		{
			`a == "aaa`,
			true,
		},
		{
			`a == "`,
			true,
		},
		{
			`a == ""'`,
			true,
		},
	}

	for _, tc := range testcases {
		expr, err := rules.ParseFromString(tc.rules)
		if tc.isError {
			assert.Error(t, err, tc.rules)
		} else {
			require.NoError(t, err)
			assert.NotNil(t, expr, tc.rules)
		}

	}
}
