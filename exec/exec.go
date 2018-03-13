package exec

import (
	"golang.org/x/net/websocket"
	"github.com/docker/docker/client"
	"fmt"
)

var cli *client.Client = createDockerClient()

type MachineExec struct {
	machineName string
	wsId string
	cmd string
	pty bool
}

func Create(ws *websocket.Conn)  {
	fmt.Println("create")
}

func Get(ws *websocket.Conn) {
	fmt.Println("get")
}

func Resize(ws *websocket.Conn)  {
	fmt.Println("resize")
}

func createDockerClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}
