package modules

import "testing"

func TestTreeOutput(t *testing.T) {
	type args struct {
		numberOfLines int
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name:     "First Test",
			args:     args{numberOfLines: 1},
			expected: "*",
		},
		{
			name: "Second Test",
			args: args{numberOfLines: 2},
			expected: "*\n" +
				" ** ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if result := TreeOutput(tt.args.numberOfLines); result != tt.expected {
				t.Errorf("TreeOutput() error = %v, expected %v", result, tt.expected)
			}
		})
	}
}
