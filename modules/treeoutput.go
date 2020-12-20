package modules

import "fmt"

//TreeOutput create a christmastree output.
func TreeOutput(numberOfLines int) (tree string) {
	switch numberOfLines {
	case 1:
		tree = "*"
	case 2:
		tree =
			" *\n" +
				"***"
	case 3:
		tree =
			"  *\n" +
				" ***\n" +
				"*****"
	case 4:
		tree =
			"   *\n" +
				"  ***\n" +
				" *****\n" +
				"*******"
	default:
		fmt.Errorf("Unknown numberOfLines given. %d", numberOfLines)
	}
	return tree
}
