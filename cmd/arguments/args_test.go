package arguments

import (
	"flag"
	"testing"
)

func TestArguments_Parse(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		expectErr  bool
		errMessage string
	}{
		{
			name:      "Valid Arguments with Algorithm",
			args:      []string{"-a", "bubble", "-s", "20", "-r"},
			expectErr: false,
		},
		{
			name:       "Missing Algorithm",
			args:       []string{"-s", "20", "-r"},
			expectErr:  true,
			errMessage: "Algorithm is required",
		},
		{
			name:       "Input and Random Together",
			args:       []string{"-a", "merge", "-i", "1,2,3,4", "-r"},
			expectErr:  true,
			errMessage: "Input and Random cannot be used together",
		},
		{
			name:      "Valid Arguments with Input",
			args:      []string{"-a", "quick", "-i", "1,2,3,4"},
			expectErr: false,
		},
		{
			name:      "Default Size Without Random",
			args:      []string{"-a", "quick"},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a custom FlagSet for this test
			flagSet := flag.NewFlagSet(tt.name, flag.ContinueOnError)

			// Initialize Arguments struct and parse the test arguments
			args := NewArguments()
			err := args.parseFlagSet(flagSet, tt.args)

			// Check for expected errors
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}

			// Validate specific error messages
			if tt.expectErr && err.Error() != tt.errMessage {
				t.Errorf("Expected error message: '%s', got: '%s'", tt.errMessage, err.Error())
			}
		})
	}
}
