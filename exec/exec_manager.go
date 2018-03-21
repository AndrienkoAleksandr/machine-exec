package exec

import (
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"errors"
	"github.com/docker/docker/api/types/filters"
	"golang.org/x/net/context"
	"sync/atomic"
)

var (
	cli = createDockerClient()
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

const (
	WsIdLabel        = "org.eclipse.che.workspace.id"
	MachineNameLabel = "org.eclipse.che.machine.name"
	FilterLabelArg   = "label"
)

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
		Cmd:          []string{machineExec.Cmd}, //todo /bin/bash -l without login ?
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

	hjr, err := cli.ContainerExecAttach(context.Background(), machineExec.ExecId, types.ExecStartCheck{})
	if err != nil {
		return nil, errors.New("Failed to attach to exec " + err.Error())
	}
	fmt.Println("attached!!!")

	return &hjr, nil
}

func Get(id string) {
	fmt.Println("get")
}

func Resize(id int, cols uint, rows uint) error {
	machineExec := execMap[id]

	resizeParam := types.ResizeOptions{Height: rows, Width:cols}
	if err := cli.ContainerExecResize(context.Background(), machineExec.ExecId, resizeParam); err != nil {
		return  err
	}
	fmt.Println("resize")
	return nil
}

func Kill() {
	//todo implement kill for exec
}

// Filter container by labels: wsId and machineName.
func findMachineContainer(identifier *model.MachineIdentifier) (*types.Container, error) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: createMachineFilter(identifier),
	})
	if err != nil {
		return nil, err
	}

	if len(containers) > 1 {
		return nil, errors.New("filter found more than one machine")
	}
	if len(containers) == 0 {
		return nil, errors.New("machine was not found")
	}

	return &containers[0], nil
}

func createMachineFilter(identifier *model.MachineIdentifier) filters.Args {
	wsIdCondition := WsIdLabel + "=" + identifier.WsId
	machineNameCondition := MachineNameLabel + "=" + identifier.MachineName

	wsfIdFilterArg := filters.Arg(FilterLabelArg, wsIdCondition)
	machineNameFilterArg := filters.Arg(FilterLabelArg, machineNameCondition)

	return filters.NewArgs(wsfIdFilterArg, machineNameFilterArg)
}
