package model

type MachineIdentifier struct {
	MachineName string
	WsId        string
}

type MachineExec struct {
	Identifier MachineIdentifier `json:"identifier"`
	Cmd        string            `json:"string"`
	Pty        bool              `json:"bool"`
	Cols       int               `json:"cols"`
	Rows       int               `json:"rows"`
	ID         int64             `json:"id"`
}
