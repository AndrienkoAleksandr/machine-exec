package exec

import (
	"errors"
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"strconv"
	"sync"
	"sync/atomic"
	"github.com/AndrienkoAleksandr/machine-exec/line-buffer"
)

type MachineExecs struct {
	mutex   *sync.Mutex
	execMap map[int]*model.MachineExec
}

var (
	cli          = createDockerClient()
	machineExecs = MachineExecs{
		mutex:   &sync.Mutex{},
		execMap: make(map[int]*model.MachineExec),
	}
	prevExecID uint64 = 0
)

// todo
//func init() {
//	body, err := cli.Events(context.Background(), types.EventsOptions{})
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	for {
//		select {
//		case err := <- err:
//			println(err)
//		case res := <- body:
//			fmt.Println(res)
//		default:
//			timer2 := time.NewTimer(5 * time.Second)
//			<-timer2.C
//			fmt.Println("default")
//		}
//	}
//}

func createDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts()//client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}

//todo don't allow to connect in the current container!!!!
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

	defer machineExecs.mutex.Unlock()
	machineExecs.mutex.Lock()

	machineExec.ExecId = resp.ID
	machineExec.ID = int(atomic.AddUint64(&prevExecID, 1))
	machineExec.Buffer = line_buffer.CreateNewLineRingBuffer()

	machineExecs.execMap[machineExec.ID] = machineExec

	fmt.Println("Create exec ", machineExec.ID, "execId", machineExec.ExecId)

	return machineExec.ID, nil
}

func Check(id int) (int, error)  {
	machineExec := getById(id)
	if &machineExec == nil {
		return -1, errors.New("Exec '" + strconv.Itoa(id) + "' was not found")
	}
	return machineExec.ID, nil
}

func Attach(id int) (*model.MachineExec, error) {
	machineExec := getById(id)
	if &machineExec == nil {
		return nil, errors.New("Exec '" + strconv.Itoa(id) + "' to attach was not found")
	}

	if machineExec.Hjr != nil {
		return machineExec, nil
	}

	hjr, err := cli.ContainerExecAttach(context.Background(), machineExec.ExecId, types.ExecStartCheck{
		Detach: false,
		Tty:    machineExec.Tty,
	})
	if err != nil {
		return nil, errors.New("Failed to attach to exec " + err.Error())
	}
	machineExec.Hjr = &hjr

	return machineExec, nil
}

func Resize(id int, cols uint, rows uint) error {
	machineExec := getById(id)
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

func getById(id int) *model.MachineExec {
	defer machineExecs.mutex.Unlock()

	machineExecs.mutex.Lock()
	return machineExecs.execMap[id]
}
