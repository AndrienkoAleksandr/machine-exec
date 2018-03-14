package exec

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/websocket"
)

var cli *client.Client = createDockerClient()

type MachineIdentier struct {
	machineName string
	wsId        string
}

type MachineExec struct {
	identifier MachineIdentier `json:"identifier"`
	cmd        string          `json:"string"`
	pty        bool            `json:"bool"`
	cols       int             `json:"cols"`
	rows       int             `json:"rows"`
	id         int64           `json:"id"`
}

func createDockerClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}

//todo exec registry ?
func Create(ws *websocket.Conn) {
	fmt.Println("create")

	machineExec, err := parseMachineExec(ws)
	if err != nil {
		fmt.Errorf(err.Error()) // todo fatality ?
		return
	}

	fmt.Println("MachineExec parsed!!!", machineExec)
	container, err := findContainer(&machineExec.identifier)
	if err != nil {
		fmt.Errorf(err.Error()) // todo fatality ?
		return
	}

	fmt.Println("found container for creation exec! id=", container.ID)
}

// todo maybe attach is better name...
func Get(ws *websocket.Conn) {
	fmt.Println("get")
}

func Resize(ws *websocket.Conn) {
	fmt.Println("resize")
}

//todo kill exec

func findContainer(identifier *MachineIdentier) (*types.Container, error) {
	return nil, nil
}

func parseMachineExec(ws *websocket.Conn) (*MachineExec, error) {
	buff := make([]byte, 8192)
	if _, err := ws.Read(buff); err != nil {
		return nil, errors.New("Failed to read machine exec body " + err.Error())
	}

	machineExec := &MachineExec{}
	if err := json.Unmarshal(buff, machineExec); err != nil {
		return nil, errors.New("Failed to parse MachineExec " + err.Error())
	}

	return machineExec, nil
}
