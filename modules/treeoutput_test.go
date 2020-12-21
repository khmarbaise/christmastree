package modules

import "testing"

/*
     *
    ***
   *****
  *******
 *********
***********
*/
func TestTreeOutput(t *testing.T) {
	type fArgs struct {
		numberOfLines int
	}
	tests := []struct {
		name     string
		args     fArgs
		expected string
	}{
		{
			name:     "First Test",
			args:     fArgs{numberOfLines: 1},
			expected: "*",
		},
		{
			name: "Second Test",
			args: fArgs{numberOfLines: 2},
			expected: " *\n" +
				"***",
		},
		{
			name: "Third Test",
			args: fArgs{numberOfLines: 3},
			expected: "  *\n" +
				" ***\n" +
				"*****",
		},
		{
			name: "Forth Test",
			args: fArgs{numberOfLines: 4},
			expected: "   *\n" +
				"  ***\n" +
				" *****\n" +
				"*******",
		},
		{
			name: "Fifth Test",
			args: fArgs{numberOfLines: 5},
			expected: "    *\n" +
				"   ***\n" +
				"  *****\n" +
				" *******\n" +
				"*********",
		},
		{
			name: "Sixth Test",
			args: fArgs{numberOfLines: 6},
			expected: "     *\n" +
				"    ***\n" +
				"   *****\n" +
				"  *******\n" +
				" *********\n" +
				"***********",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if result := TreeOutput(tt.args.numberOfLines); result != tt.expected {
				t.Errorf("TreeOutput() error = %v, expected '%v'", result, tt.expected)
			}
		})
	}
}
