package main

import (
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	execManager "github.com/AndrienkoAleksandr/machine-exec/exec"
	"io"
	"net"
	"os"
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

	writeToExec(hiJackRepsp.Conn)

	go func() {
		io.Copy(os.Stdout, hiJackRepsp.Reader)
	}()

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	hiJackRepsp.Conn.Close()
	hiJackRepsp.Close()
	hiJackRepsp.CloseWrite()
}

func writeToExec(hjrConn net.Conn) {
	if _, err := hjrConn.Write([]byte("ls -a\n")); err != nil {
		fmt.Println(err.Error())
	}
}
