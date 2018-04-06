package main

import (
"fmt"
"github.com/AndrienkoAleksandr/machine-exec/api/model"
execManager "github.com/AndrienkoAleksandr/machine-exec/exec"
"time"
	"io"
	"os"
)

func main() {
	machineExec := model.MachineExec{
		Identifier: model.MachineIdentifier{
			MachineName: "theia",
			WsId:        "workspace8xa30590jdfzi4gb",
		},
		//Cmd:  []string{"bash", "-c", "echo a & echo $!"},
		//Cmd:  []string{"bash", "-c", "for i in {1..5}; do echo a; done & echo $!"},
		//Cmd:  []string{"bash", "-c", "for i in {1..5}; do echo \"a\" && echo 'b'; done"},
		//Cmd:  []string{"bash", "-c", "echo process1 && cat /home/theia/package.json"},
		//Cmd:  []string{"bash", "-c", "echo u1 > /dev/null && bash"},
		Cmd:  []string{"bash"},
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


	//printPid(hiJackRepsp.Reader)
	//READ(hiJackRepsp.Reader)


	go func() {
		fmt.Println("--------------------------")
		io.Copy(os.Stdout, hiJackRepsp.Reader)
		fmt.Println("--------------------------")
	}()

	timer1 := time.NewTimer(11 * time.Second)
	<-timer1.C
	hiJackRepsp.Conn.Close()
	hiJackRepsp.Close()
	hiJackRepsp.CloseWrite()
}


