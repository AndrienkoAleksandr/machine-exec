package main

import (
	execManager "github.com/AndrienkoAleksandr/machine-exec/exec"
	"fmt"
)

func main() {
	if res, ok := execManager.ExecIsAlive("18f17ffb3b1de20a7e120fd300510331cd7c1560514adc890932c46506df1d23"); ok {
		fmt.Print(res)
	}
}