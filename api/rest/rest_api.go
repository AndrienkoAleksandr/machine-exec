package rest

import "github.com/eclipse/che/agents/go-agents/core/rest"

var HTTPRoutes = rest.RoutesGroup{
	Name: "Rest MachineExec Routes",
	Items: []rest.Route{
		{
			Method:     "POST",
			Name:       "Create MachineExec",
			Path:       "/machine-exec",
			HandleFunc: CreateExec,
		},
		{
			Method:     "DELETE",
			Name:       "Kill MachineExec",
			Path:       "/machine-exec/:pid",
			HandleFunc: KillExec,
		},
		{
			Method:     "GET",
			Name:       "Get MachineExec info",
			Path:       "/machine-exec/:id",
			HandleFunc: GetExec,
		},
		{
			Method:     "GET",
			Name:       "Resize MachineExec",
			Path:       "/machine-exec/:id/resize",
			HandleFunc: ResizeExec,
		},
	},
}
