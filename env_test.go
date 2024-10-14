// SPDX-FileCopyrightText: Copyright (c) 2024, Obviously, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-obvious/env"
)

func TestGetArray(t *testing.T) {
	tests := []struct {
		name        string
		envValue    string
		sep         string
		allowQuotes bool
		trimQuoted  bool
		expected    []string
	}{
		{
			name:        "simple comma separated values",
			envValue:    "val1,val2,val3",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    []string{"val1", "val2", "val3"},
		},
		{
			name:        "values with spaces",
			envValue:    " val1 , val2 , val3 ",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    []string{"val1", "val2", "val3"},
		},
		{
			name:        "empty array",
			envValue:    "",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    nil,
		},
		{
			name:        "array with brackets",
			envValue:    "[val1,val2,val3]",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    []string{"val1", "val2", "val3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("TEST_ENV", tt.envValue)
			defer os.Unsetenv("TEST_ENV")

			result := env.GetArray("TEST_ENV", tt.sep, tt.allowQuotes, tt.trimQuoted)
			assert.Equal(t, tt.expected, result)
		})
	}
}
