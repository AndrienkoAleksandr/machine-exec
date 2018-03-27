package exec

import (
	"errors"
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"strconv"
	"sync/atomic"
	"syscall"
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
		Detach:       false,           //todo support detach exec ? Maybe for kill it would be nice...
		Cmd:          machineExec.Cmd, // todo /bin/bash -l without login ?
	})
	if err != nil {
		return -1, err
	}

	machineExec.ExecId = resp.ID

	//generate unique id
	id := int(atomic.AddUint64(&prevExecID, 1))
	execMap[id] = *machineExec

	fmt.Println("Create exec ", machineExec.ID)

	return id, nil
}

func Attach(id int) (*types.HijackedResponse, error) {
	machineExec := execMap[id]
	if &machineExec == nil {
		return nil, errors.New("Exec '" + strconv.Itoa(id) + "' to attach was not found")
	}

	hjr, err := cli.ContainerExecAttach(context.Background(), machineExec.ExecId, types.ExecStartCheck{
		Detach: false, //todo support detach exec ? Maybe for kill it would be nice...
		Tty:    machineExec.Tty,
	})
	if err != nil {
		return nil, errors.New("Failed to attach to exec " + err.Error())
	}
	fmt.Println("attach to exec")

	return &hjr, nil
}

func Get(id int) (*model.MachineExec, error) {
	machineExec := execMap[id]
	if &machineExec == nil {
		return nil, errors.New("Exec '" + strconv.Itoa(id) + "' was not found")
	}

	fmt.Println("get exec info")

	return &machineExec, nil
}

func Resize(id int, cols uint, rows uint) error {
	machineExec := execMap[id]
	if &machineExec == nil {
		return errors.New("Exec to resize '" + strconv.Itoa(id) + "' was not found")
	}

	resizeParam := types.ResizeOptions{Height: rows, Width: cols}
	if err := cli.ContainerExecResize(context.Background(), machineExec.ExecId, resizeParam); err != nil {
		return err
	}

	fmt.Println("resize")

	return nil
}

func Kill(id int) error {
	machineExec := execMap[id]
	if &machineExec == nil {
		return errors.New("Exec to kill '" + strconv.Itoa(id) + "' was not found")
	}

	execInspect, err := cli.ContainerExecInspect(context.Background(), machineExec.ExecId)
	if err != nil {
		return err
	}

	if !execInspect.Running {
		return errors.New("Exec with id '" + strconv.Itoa(id) + "' has already terminated")
	}

	pid := execInspect.Pid
	if err := syscall.Kill(-pid, syscall.SIGHUP); err != nil {
		return err
	}
	// todo send SIGHUP for pty execs, and sent SIGINT for all another execs.
	//killCommand := []string{"kill", "-" + syscall.SIGHUP.String(), strconv.Itoa(pid)}
	//killMachineExec := model.MachineExec{
	//	Identifier: model.MachineIdentifier{
	//		WsId:        machineExec.Identifier.WsId,
	//		MachineName: machineExec.Identifier.MachineName,
	//	},
	//	Tty: false,
	//	Cmd: killCommand,
	//}

	// _, err = Create(&killMachineExec)
	//if err != nil {
	//	return err
	//}
	// todo check if exec was really terminated, if not => kill by SIGKILL...

	return nil
}
