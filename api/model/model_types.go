package model

import (
	"github.com/docker/docker/api/types"
	"github.com/AndrienkoAleksandr/machine-exec/line-buffer"
)

//todo remove workspace id
type MachineIdentifier struct {
	MachineName string `json:"machineName"`
	WsId        string `json:"workspaceId"`
}

type MachineExec struct {
	Identifier MachineIdentifier `json:"identifier"`
	Cmd        []string          `json:"cmd"`
	Tty        bool              `json:"tty"`
	Cols       int               `json:"cols"`
	Rows       int               `json:"rows"`

	// unique client id, real execId should be hidden from client to prevent serialization
	ID     int `json:"id"`
	ExecId string
	Hjr *types.HijackedResponse
	Buffer line_buffer.LineRingBuffer
}
