package jsonrpc

import (
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	execManager "github.com/AndrienkoAleksandr/machine-exec/exec"
	"github.com/eclipse/che/agents/go-agents/core/jsonrpc"
	"strconv"
)

type IdParam struct {
	Id int `json:"id"` //todo maybe string like id, or int64...
}

type OperationResult struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type ResizeParam struct {
	Id   int  `json:"id"`
	Cols uint `json:"cols"`
	Rows uint `json:"rows"`
}

//TODO improvements: todo check casting param is ok...
func jsonRpcCreateExec(_ *jsonrpc.Tunnel, params interface{}, t jsonrpc.RespTransmitter) {
	machineExec := params.(*model.MachineExec)

	execId, err := execManager.Create(machineExec)
	if err != nil {
		t.SendError(jsonrpc.NewArgsError(err))
	}

	t.Send(execId)
}

func jsonRpcResizeExec(_ *jsonrpc.Tunnel, params interface{}) (interface{}, error) {
	resizeParam := params.(*ResizeParam)

	if err := execManager.Resize(resizeParam.Id, resizeParam.Cols, resizeParam.Rows); err != nil {
		return nil, jsonrpc.NewArgsError(err) //todo as jsonRpc error?
	}
	fmt.Println("Resize with json RPC!")

	return &OperationResult{
		Id: resizeParam.Id, Text: "Exec with id " + strconv.Itoa(resizeParam.Id) + "  was successfully resized",
	}, nil
}

func jsonRpcKillExec(_ *jsonrpc.Tunnel, params interface{}) (interface{}, error) {
	idParam := params.(*IdParam)

	err := execManager.Kill(idParam.Id)
	if err != nil {
		return nil, err
	}

	fmt.Println("Kill with json RPC!")

	return &OperationResult{
		Id:   idParam.Id,
		Text: "Exec with id '" + strconv.Itoa(idParam.Id) + "' was successfully killed",
	}, nil
}
