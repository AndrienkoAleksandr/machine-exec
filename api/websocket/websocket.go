package websocket

import (
	"fmt"
	"log"
	"github.com/eclipse/che-lib/websocket"
	"github.com/eclipse/che/agents/go-agents/core/rest"
	"net/http"
	execManager "github.com/AndrienkoAleksandr/machine-exec/exec"
	//"strconv"
	//"errors"
	"strconv"
	"errors"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//todo apply ping-pong handler!!!
)

func Attach(w http.ResponseWriter, r *http.Request, restParmas rest.Params) error {
	id, err := strconv.Atoi(restParmas.Get("id"));
	if err != nil {
		return errors.New("Failed to parse id")
	}
	fmt.Println("Parsed id", id)

	hjr, err := execManager.Attach(id)
	if err != nil {
		return err
	}

	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error " + err.Error())
		return err
	}

	//todo implement read/write goroutines for websocket connection

	fmt.Println("Connection!", con != nil, hjr != nil)
	return nil
}
