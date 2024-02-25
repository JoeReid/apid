package main

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	var tests = []struct {
		name         string
		stdin        string
		args         []string
		expectOutput string
		checkErr     func(error) bool
	}{
		{
			name:         "encode arg no prefix",
			stdin:        "",
			args:         []string{"e8ce617a-4f10-4b3b-9748-d4c34c0c77b4"},
			expectOutput: "unknown_75ixeUANbpM4NtrDZs1GzW",
			checkErr:     noError,
		},
		{
			name:         "encode arg with prefix",
			stdin:        "",
			args:         []string{"-p", "test", "e8ce617a-4f10-4b3b-9748-d4c34c0c77b4"},
			expectOutput: "test_75ixeUANbpM4NtrDZs1GzW",
			checkErr:     noError,
		},
		{
			name:         "endocde arg with invalid uuid",
			stdin:        "",
			args:         []string{"-p", "test", "e8ce617a-4f10-4b3b-9748-d4c34c0c77b4-"},
			expectOutput: "",
			checkErr:     errContains("invalid UUID"),
		},
		{
			name:         "encode stdin no prefix",
			stdin:        "e8ce617a-4f10-4b3b-9748-d4c34c0c77b4",
			args:         []string{},
			expectOutput: "unknown_75ixeUANbpM4NtrDZs1GzW",
			checkErr:     noError,
		},
		{
			name:         "encode stdin with prefix",
			stdin:        "e8ce617a-4f10-4b3b-9748-d4c34c0c77b4",
			args:         []string{"-p", "test"},
			expectOutput: "test_75ixeUANbpM4NtrDZs1GzW",
			checkErr:     noError,
		},
		{
			name:         "encode stdin with invalid uuid",
			stdin:        "e8ce617a-4f10-4b3b-9748-d4c34c0c77b4-",
			args:         []string{"-p", "test"},
			expectOutput: "",
			checkErr:     errContains("invalid UUID"),
		},
		{
			name:         "decode arg",
			stdin:        "",
			args:         []string{"-d", "test_75ixeUANbpM4NtrDZs1GzW"},
			expectOutput: "e8ce617a-4f10-4b3b-9748-d4c34c0c77b4",
			checkErr:     noError,
		},
		{
			name:         "decode arg with invalid apid",
			stdin:        "",
			args:         []string{"-d", "test_75ixeUANbpM4NtrDZs1GzW-"},
			expectOutput: "",
			checkErr:     errContains("invalid APID"),
		},
		{
			name:         "decode stdin",
			stdin:        "test_75ixeUANbpM4NtrDZs1GzW",
			args:         []string{"-d"},
			expectOutput: "e8ce617a-4f10-4b3b-9748-d4c34c0c77b4",
			checkErr:     noError,
		},
		{
			name:         "decode stdin with invalid apid",
			stdin:        "test_75ixeUANbpM4NtrDZs1GzW-",
			args:         []string{"-d"},
			expectOutput: "",
			checkErr:     errContains("invalid APID"),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			candidate := CLI{
				Stdin: strings.NewReader(tt.stdin),
				Args:  tt.args,
			}

			output, err := candidate.Run()
			if !tt.checkErr(err) {
				t.Errorf("unexpected error: %v", err)
			}

			if output != tt.expectOutput {
				t.Errorf("unexpected output: got %q, want %q", output, tt.expectOutput)
			}
		})
	}
}

func noError(err error) bool {
	return err == nil
}

func errContains(substr string) func(error) bool {
	return func(err error) bool {
		return err != nil && strings.Contains(err.Error(), substr)
	}
}
