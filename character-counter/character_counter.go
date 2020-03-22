package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	line = bytes.TrimRight(line, "\r\n")
	input := string(line)
	// fmt.Println(input)

	fmt.Println("number of bytes", len(input))
	fmt.Println("number of runes", utf8.RuneCountInString(input))

	letterCounter, spaceCounter, digitCounter, chineseCounter, otherCounter := 0, 0, 0, 0, 0
	for _, r := range input {
		// fmt.Printf("%d %v %T %s\n", i, r, r, string(r))

		if 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' {
			letterCounter++
		} else if unicode.IsSpace(r) {
			spaceCounter++
		} else if unicode.IsDigit(r) {
			digitCounter++
		} else if unicode.Is(unicode.Scripts["Han"], r) {
			chineseCounter++
		} else {
			otherCounter++
		}

	}

	total := letterCounter + spaceCounter + digitCounter + chineseCounter + otherCounter
	if total != utf8.RuneCountInString(input) {
		panic("error")
	}

	fmt.Printf("total %d, letter %d, space %d, digit %d, chinese %d, other %d\n", total, letterCounter, spaceCounter, digitCounter, chineseCounter, otherCounter)
}
