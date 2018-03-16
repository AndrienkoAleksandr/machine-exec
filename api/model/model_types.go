package model

type MachineIdentifier struct {
	MachineName string `json:"machineName"`
	WsId        string `json:"workspaceId"`
}

type MachineExec struct {
	Identifier MachineIdentifier `json:"identifier"`
	Cmd        string            `json:"cmd"`
	Pty        bool              `json:"pty"`
	Cols       int               `json:"cols"`
	Rows       int               `json:"rows"`
	ID         int64             `json:"id"`
}
