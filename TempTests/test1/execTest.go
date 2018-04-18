package main

import (
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	execManager "github.com/AndrienkoAleksandr/machine-exec/exec"
	"net"
	"time"
	"io"
	"os"
)

func main() {
	machineExec := model.MachineExec{
		Identifier: model.MachineIdentifier{
			MachineName: "dev",
			WsId:        "workspacecs82k5zp6jyv86fs",
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
	machineExecFilled, err := execManager.Attach(id)
	if err != nil {
		fmt.Println(err.Error())
	}

	hjr := machineExecFilled.Hjr


	writeToExec(hjr.Conn)

	go func() {
		io.Copy(os.Stdout, hjr.Reader)
	}()


	timer1 := time.NewTimer(2 * time.Second)
	bst, _, _  := hjr.Reader.ReadLine()
	fmt.Println(string(bst))

	<-timer1.C
	timer2 := time.NewTimer(2 * time.Second)

	bst2, _, _  := hjr.Reader.ReadLine()
	fmt.Println(string(bst2))

	<-timer2.C
	hjr.Conn.Close()
	hjr.Close()
	hjr.CloseWrite()
}

func writeToExec(hjrConn net.Conn) {
	if _, err := hjrConn.Write([]byte("ls -a\n")); err != nil {
		fmt.Println(err.Error())
	}
}
