package exec

import (
	"errors"
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"sync/atomic"
)

var (
	cli               = createDockerClient()
	execMap           = make(map[int]model.MachineExec) // todo think about multi-threads!!!
	prevExecID uint64 = 0
)

func createDockerClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}

func Create(machineExec *model.MachineExec) (int, error) {
	container, err := findMachineContainer(&machineExec.Identifier)
	if err != nil {
		return -1, err
	}

	fmt.Println("found container for creation exec! id=", container.ID)

	resp, err := cli.ContainerExecCreate(context.Background(), container.ID, types.ExecConfig{
		Tty:          machineExec.Tty,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Detach:       false,                     //todo support detach exec ? Maybe for kill it would be nice...
		Cmd:          []string{machineExec.Cmd}, // todo /bin/bash -l without login ?
	})
	if err != nil {
		return -1, err
	}

	machineExec.ExecId = resp.ID

	//generate unique id
	id := int(atomic.AddUint64(&prevExecID, 1))
	execMap[id] = *machineExec

	return id, nil
}

func Attach(id int) (*types.HijackedResponse, error) {
	machineExec := execMap[id]

	hjr, err := cli.ContainerExecAttach(context.Background(), machineExec.ExecId, types.ExecStartCheck{
		Detach: false, //todo support detach exec ? Maybe for kill it would be nice...
		Tty:    machineExec.Tty,
	})
	if err != nil {
		return nil, errors.New("Failed to attach to exec " + err.Error())
	}
	fmt.Println("attached!!!")

	return &hjr, nil
}

func Get(id string) {
	// todo implement method get
	fmt.Println("get")
}

func Resize(id int, cols uint, rows uint) error {
	machineExec := execMap[id]

	resizeParam := types.ResizeOptions{Height: rows, Width: cols}
	if err := cli.ContainerExecResize(context.Background(), machineExec.ExecId, resizeParam); err != nil {
		return err
	}
	fmt.Println("resize")
	return nil
}

func Kill() {
	//todo implement kill for exec
}
