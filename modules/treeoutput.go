package modules

import (
	"fmt"
	"strings"
)

//TreeOutput create a christmastree output.
func TreeOutput(numberOfLines int) (tree string) {

	switch numberOfLines {
	case 1:
		tree =
			strings.Repeat("*", numberOfLines*2-1)
	case 2:
		tree =
			" *\n" +
				strings.Repeat("*", numberOfLines*2-1)
	case 3:
		tree =
			"  *\n" +
				" ***\n" +
				strings.Repeat("*", numberOfLines*2-1)
	case 4:
		tree =
			"   *\n" +
				"  ***\n" +
				" *****\n" +
				strings.Repeat("*", numberOfLines*2-1)
	case 5:
		tree =
			"    *\n" +
				"   ***\n" +
				"  *****\n" +
				" *******\n" +
				strings.Repeat("*", numberOfLines*2-1)
	default:
		fmt.Errorf("Unknown numberOfLines given. %d", numberOfLines)
	}
	return
}
