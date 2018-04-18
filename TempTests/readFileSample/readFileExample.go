package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fs, err := os.Open("TempTests/readFileSample/readFileExample.go")

	if err != nil {
		fmt.Println("error reading file", err.Error())
		return
	}

	defer fs.Close()

	reader := bufio.NewReader(fs)

	buf := make([]byte, 100)

	for {
		v, _ := reader.Read(buf)

		if v == 0 {
			return
		}

		fmt.Print(string(buf[0:v]))
	}
}
