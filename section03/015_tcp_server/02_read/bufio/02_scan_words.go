package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "When the praises go up,\nThe blessings come down.\nWhen the praises go up,\nThe blessings come down.\nIt seems like blessings\nKeep falling in my lap.\nIt seems like blessings\nKeep falling in my lap.\n"

	scanner := bufio.NewScanner(strings.NewReader(s))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
