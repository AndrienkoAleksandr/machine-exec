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
	Id   int64  `json:"id"`
	Text string `json:"text"`
}

//todo implement it
func jsonRpcResizeExec(_ *jsonrpc.Tunnel, params interface{}) (interface{}, error) {
	//machine := params.(*model.MachineExec)
	//if err := machineManager.resize(machine); err != nil {
	//	return nil, asRPCError(err)
	//}
	fmt.Println("Resize with json RPC!")

	return &OperationResult{Id: 123, Text: "Successfully resize"}, nil
}

type KillParam struct {
	Id   int  `json:"id"`
}

//todo implement it
func jsonRpcKillExec(_ *jsonrpc.Tunnel, params interface{}) (interface{}, error) {
	fmt.Println("Kill with json RPC!")
	return nil, nil
}
