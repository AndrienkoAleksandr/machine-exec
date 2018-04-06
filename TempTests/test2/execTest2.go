package main

import (
	"bufio"
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	execManager "github.com/AndrienkoAleksandr/machine-exec/exec"
	"net"
	"time"
)

func main() {
	machineExec := model.MachineExec{
		Identifier: model.MachineIdentifier{
			MachineName: "dev-machine",
			WsId:        "workspacemru4loxoylowd537",
		},
		Cmd:  []string{"/bin/bash"},
		Cols: 24,
		Rows: 80,
		Tty:  true,
	}
	id, err := execManager.Create(&machineExec)
	if err != nil {
		fmt.Println(err.Error())
	}
	hiJackRepsp, err := execManager.Attach(id)
	if err != nil {
		fmt.Println(err.Error())
	}

	readAndPrint(hiJackRepsp.Reader)
	writeToExec(hiJackRepsp.Conn)

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	readAndPrint(hiJackRepsp.Reader)
}

func readAndPrint(reader *bufio.Reader) {
	bts := make([]byte, 8192)
	size, err := reader.Read(bts)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bts[0:size]))
}

func writeToExec(hjrConn net.Conn) {
	if _, err := hjrConn.Write([]byte("ls -a\n")); err != nil {
		fmt.Println(err.Error())
	}
}
