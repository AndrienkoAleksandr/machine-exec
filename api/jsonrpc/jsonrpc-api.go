package jsonrpc

import (
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	"github.com/eclipse/che/agents/go-agents/core/jsonrpc"
)

// Constants that represent RPC methods identifiers.
const (
	CreateMethod = "create.exec"
	GetMethod    = "get.exec"
	ResizeMethod = "resize.exec"
)

// Error codes.
const (
	ProcessAPIErrorCode      = 100
	NoSuchProcessErrorCode   = 101
	ProcessNotAliveErrorCode = 102
)

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
			Decode: jsonrpc.FactoryDec(func() interface{} { return &OperationResult{} }),
			Handle: jsonrpc.HandleRet(jsonrpcResizeExec),
		},
	},
}
