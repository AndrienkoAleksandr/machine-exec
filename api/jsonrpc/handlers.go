package jsonrpc

import (
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	execManager "github.com/AndrienkoAleksandr/machine-exec/exec"
	"github.com/eclipse/che/agents/go-agents/core/jsonrpc"
)

func jsonRpcCreateExec(_ *jsonrpc.Tunnel, params interface{}, t jsonrpc.RespTransmitter) {
	machineExec := params.(*model.MachineExec)

	execId, err := execManager.Create(machineExec)
	if err != nil {
		t.SendError(jsonrpc.NewArgsError(err))
	}

	t.Send(execId)
}

func jsonRpcGetExec(_ *jsonrpc.Tunnel, params interface{}, t jsonrpc.RespTransmitter) {
	machineExec := params.(*model.MachineExec)

	fmt.Println("Get with json RPC!")

	t.Send(machineExec)
}

type OperationResult struct {
	Id   int    `json:"id"` //todo maybe string like id, or int64...
	Text string `json:"text"`
}

type ResizeParam struct {
	Id   int  `json:"id"`
	Cols uint `json:"cols"`
	Rows uint `json:"rows"`
}

func jsonRpcResizeExec(_ *jsonrpc.Tunnel, params interface{}) (interface{}, error) {
	resizeParam := params.(*ResizeParam)

	if err := execManager.Resize(resizeParam.Id, resizeParam.Cols, resizeParam.Rows); err != nil {
		//todo as jsonRpc error?
		return nil, jsonrpc.NewArgsError(err)
	}
	fmt.Println("Resize with json RPC!")

	return &OperationResult{Id: 123, Text: "Successfully resize"}, nil
}

type KillParam struct {
	Id int `json:"id"`
}

//todo implement it
func jsonRpcKillExec(_ *jsonrpc.Tunnel, params interface{}) (interface{}, error) {
	execManager.Kill();
	fmt.Println("Kill with json RPC!")
	return nil, nil
}
