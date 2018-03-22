package readFileSample

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fs, err := os.Open("/home/user/GoWorkSpace/src/github.com/AndrienkoAleksandr/machine-exec/api/websocket/websocket.go")

	if err != nil {
		fmt.Println("error reading file")
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
