package jsonrpc

import (
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	"github.com/eclipse/che/agents/go-agents/core/jsonrpc"
)

// Constants that represent RPC methods identifiers.
const (
	CreateMethod = "create"
	GetMethod    = "get"
	ResizeMethod = "resize"
	KillMethod   = "kill"
)

// Error codes.
const (
	ProcessAPIErrorCode      = 100
	NoSuchProcessErrorCode   = 101
	ProcessNotAliveErrorCode = 102
)

//TODO we need check casting for all param types!!!

// RPCRoutes defines process jsonrpc routes.
var RPCRoutes = jsonrpc.RoutesGroup{
	Name: "Json-rpc MachineExec Routes",
	Items: []jsonrpc.Route{
		{
			Method: CreateMethod,
			Decode: jsonrpc.FactoryDec(func() interface{} { return &model.MachineExec{} }),
			Handle: jsonRpcCreateExec,
		},
		{
			Method: GetMethod,
			Decode: jsonrpc.FactoryDec(func() interface{} { return &model.MachineExec{} }),
			Handle: jsonRpcGetExec,
		},
		{
			Method: ResizeMethod,
			Decode: jsonrpc.FactoryDec(func() interface{} { return &ResizeParam{} }),
			Handle: jsonrpc.HandleRet(jsonRpcResizeExec),
		},
		{
			Method: KillMethod,
			Decode: jsonrpc.FactoryDec(func() interface{} { return &OperationResult{} }),
			Handle: jsonrpc.HandleRet(jsonRpcKillExec),
		},
	},
}

// create:
// {"jsonrpc":"2.0","method":"create", "id":"12","params":{"identifier":{"machineName":"dev-machine", "workspaceId":"workspacethvmvurngk7sjrvb"}, "cmd":"/bin/bash", "pty": true, "cols":24, "rows":80, "id": 900 }}
