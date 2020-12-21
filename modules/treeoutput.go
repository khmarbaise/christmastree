package modules

import (
	"strings"
)

//TreeOutput create a christmastree output.
func TreeOutput(numberOfLines int) (tree string) {

	for i := 1; i < numberOfLines; i++ {
		tree += strings.Repeat(" ", numberOfLines-i) + strings.Repeat("*", (i*2)-1) + "\n"
	}

	tree += strings.Repeat("*", numberOfLines*2-1)
	return
}
