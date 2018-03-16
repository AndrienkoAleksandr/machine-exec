package main

import (
	"flag"
	jsonRpcApi "github.com/AndrienkoAleksandr/machine-exec/api/jsonrpc"
	restApi "github.com/AndrienkoAleksandr/machine-exec/api/rest"
	"log"
	"net/http"
	"time"
	//"golang.org/x/net/websocket"
	"github.com/eclipse/che/agents/go-agents/core/jsonrpc"
	"github.com/eclipse/che/agents/go-agents/core/jsonrpc/jsonrpcws"
	"github.com/eclipse/che/agents/go-agents/core/rest"
	"github.com/eclipse/che-lib/websocket"
)

var (
	url, filesPath string

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//todo ping-pong handler
)

func init() {
	flag.StringVar(&url, "url", ":4444", "Host:Port address. ")
	flag.StringVar(&filesPath, "client", "./client", "Path to the files to serve them.")
}

func main() {
	flag.Parse()

	//todo apply exec-machine context again

	appRoutes := []rest.RoutesGroup{
		restApi.HTTPRoutes,
		{
			Name: "Exec-Machine WebSocket routes",
			Items: []rest.Route{
				{
					Method: "GET",
					Path:   "/connect",
					Name:   "Connect to MachineExec api(websocket)",
					HandleFunc: func(w http.ResponseWriter, r *http.Request, _ rest.Params) error {
						conn, err := jsonrpcws.Upgrade(w, r)
						if err != nil {
							return err
						}
						tunnel := jsonrpc.NewManagedTunnel(conn)
						tunnel.SayHello()
						return nil
					},
				},
				{
					Method: "GET",
					Path:   "/pty", //todo apply id!!!!!
					Name:   "Connect to pty exec(websocket)",
					HandleFunc: func(w http.ResponseWriter, r *http.Request, _ rest.Params) error {
						log.Println("Connect to gorilla pty")
						conn, err := upgrader.Upgrade(w, r, nil)
						if err != nil {
							log.Println("Error " + err.Error())
							return err
						}

						defer conn.Close()

						conn.WriteMessage(websocket.TextMessage, []byte("Hello from websocket for pty"))
						if err != nil {
							log.Println("Error " + err.Error())
							return err
						}
						return nil
					},
				},
			},
		},
	}

	// create json-rpc routs group
	appOpRoutes := []jsonrpc.RoutesGroup{
		jsonRpcApi.RPCRoutes,
	}

	// register routes and http handlers
	baseHandler := rest.NewDefaultRouter(url, appRoutes)
	rest.PrintRoutes(appRoutes)
	jsonrpc.RegRoutesGroups(appOpRoutes)
	jsonrpc.PrintRoutes(appOpRoutes)

	server := &http.Server{
		Handler:      baseHandler,
		Addr:         url,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	server.ListenAndServe()
}
