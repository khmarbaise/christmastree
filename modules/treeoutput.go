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
			strings.Repeat(" ", numberOfLines-1) + "*\n"
	case 3:
		tree =
			strings.Repeat(" ", numberOfLines-1) + strings.Repeat("*", 1) + "\n" +
				strings.Repeat(" ", numberOfLines-2) + strings.Repeat("*", 3) + "\n"
	case 4:
		tree =
			strings.Repeat(" ", numberOfLines-1) + strings.Repeat("*", 1) + "\n" +
				strings.Repeat(" ", numberOfLines-2) + strings.Repeat("*", 3) + "\n" +
				strings.Repeat(" ", numberOfLines-3) + strings.Repeat("*", 5) + "\n"
	case 5:
		tree =
			strings.Repeat(" ", numberOfLines-1) + strings.Repeat("*", 1) + "\n" +
				strings.Repeat(" ", numberOfLines-2) + strings.Repeat("*", 3) + "\n" +
				strings.Repeat(" ", numberOfLines-3) + strings.Repeat("*", 5) + "\n" +
				strings.Repeat(" ", numberOfLines-4) + strings.Repeat("*", 7) + "\n"
	case 6:
		tree =
			strings.Repeat(" ", numberOfLines-1) + strings.Repeat("*", 1) + "\n" +
				strings.Repeat(" ", numberOfLines-2) + strings.Repeat("*", 3) + "\n" +
				strings.Repeat(" ", numberOfLines-3) + strings.Repeat("*", 5) + "\n" +
				strings.Repeat(" ", numberOfLines-4) + strings.Repeat("*", 7) + "\n" +
				strings.Repeat(" ", numberOfLines-5) + strings.Repeat("*", 9) + "\n"
	default:
		fmt.Errorf("Unknown numberOfLines given. %d", numberOfLines)
	}

	tree += strings.Repeat("*", numberOfLines*2-1)
	return
}
