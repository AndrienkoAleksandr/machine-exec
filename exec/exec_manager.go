package exec

import (
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	//"golang.org/x/net/websocket"
)

var cli *client.Client = createDockerClient()

func createDockerClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}

//func Attach(ws *websocket.Conn) {
//
//}

//todo exec registry ?
func Create(machineExec model.MachineExec) {
	fmt.Println("create")

	fmt.Println("MachineExec parsed!!!", machineExec)
	container, err := findContainer(&machineExec.Identifier)
	if err != nil {
		fmt.Errorf(err.Error()) // todo fatality ?
		return
	}

	//todo cli create logic

	fmt.Println("found container for creation exec! id=", container.ID)
}

func Get(id string) {
	fmt.Println("get")
}

func Resize(id string) {
	fmt.Println("resize")
}

//todo kill exec

func findContainer(identifier *model.MachineIdentifier) (*types.Container, error) {
	return nil, nil
}
