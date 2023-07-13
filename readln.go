package bcfl
// BCF Libs

import (
	"bufio"
	"os"
	"strings"
)

/*
 * Readln(): Read from stdin and return it as a string
 */
func Readln() string {
	var reader = bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// Strip final '\n'.
	input = strings.Trim(input, "\n")

	return input
}
