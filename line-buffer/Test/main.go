package main

import (
	"bufio"
	"strings"
	"fmt"
	"bytes"
)

func main()  {
	//bufio.ScanRunes()
	//bufio.

	//input := "foo   bar      baz\ngg"
	input := "\nz\ngg\n\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	//scanner.Split(bufio.ScanLines)
	scanner.Split(ScanLinesNoDropCR)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>")
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>")


}

func ScanLinesNoDropCR(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i + 1], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
