package jsonrpc

import (
	"fmt"
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	"github.com/eclipse/che/agents/go-agents/core/jsonrpc"
	"log"
)

type OperationResult struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
}

func jsonrpcCreateExec(tun *jsonrpc.Tunnel, params interface{}, t jsonrpc.RespTransmitter) {
	machineExec := params.(*model.MachineExec)

	fmt.Println("Create with json RPC!")

	//t.SendError(jsonrpc.NewArgsError(errors.New("Something went wrong")));

	t.Send(machineExec)
	//tun.Notify("result", machineExec)
}

func jsonrpcGetExec(_ *jsonrpc.Tunnel, params interface{}, t jsonrpc.RespTransmitter) {
	machineExec := params.(*model.MachineExec)

	fmt.Println("Get with json RPC!")

	//t.SendError(jsonrpc.NewArgsError(errors.New("Something went wrong")));
	t.Send(machineExec)
}

func jsonrpcResizeExec(_ *jsonrpc.Tunnel, params interface{}) (interface{}, error) {
	//machine := params.(*model.MachineExec)
	//if err := machineManager.resize(machine); err != nil {
	//	return nil, asRPCError(err)
	//}
	fmt.Println("Resize with json RPC!")

	return &OperationResult{Id: 123, Text: "Successfully resize"}, nil
}


func jsonrpcStartExec(tun *jsonrpc.Tunnel, params interface{}, t jsonrpc.RespTransmitter) {
	machineExec := params.(*model.MachineExec)

	t.Send(machineExec)
}
