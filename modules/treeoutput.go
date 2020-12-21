package modules

import (
	"fmt"
	"strings"
)

//TreeOutput create a christmastree output.
func TreeOutput(numberOfLines int) (tree string) {

	switch numberOfLines {
	case 1:
		tree = ""
	case 2:
		tree =
			" *\n"
	case 3:
		tree =
			"  *\n" +
				" ***\n"
	case 4:
		tree =
			"   *\n" +
				"  ***\n" +
				" *****\n"
	case 5:
		tree =
			"    *\n" +
				"   ***\n" +
				"  *****\n" +
				" *******\n"
	default:
		fmt.Errorf("Unknown numberOfLines given. %d", numberOfLines)
	}
	tree += strings.Repeat("*", numberOfLines*2-1)
	return
}
