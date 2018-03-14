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
		//todo
		//{
		//	Method:     "DELETE",
		//	Name:       "Kill Exec",
		//	Path:       "/process/:pid",
		//	HandleFunc: detach,
		//},
		{
			Method:     "GET",
			Name:       "Get MachineExec info",
			Path:       "/machine-exec/:id",
			HandleFunc: GetExec,
		},

		// todo think about update method... but for now we can change only exec geometry...
		// So let be separated resize method for now...
		{
			Method:     "GET",
			Name:       "Resize MachineExec",
			Path:       "/machine-exec/:id/resize",
			HandleFunc: ResizeExec,
		},
	},
}
