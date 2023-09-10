package bcfl

import (
	"bufio"
	"fmt"
	"os"
)

func FileToSlice (fpath string) []string {
	file, err := os.Open(fpath)

	if err != nil {
		fmt.Println(err)
	}

	var lines []string

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s := fileScanner.Text()
		lines = append(lines, s)
	}

	file.Close()

	return lines
}

/*
func main() {
	arq := f2slice("arq.txt")
	//fmt.Printf("%v", arq)
	for i, s := range arq {
		fmt.Printf("%d: %s\n", i, s)
	}
	fmt.Println()
}
*/

