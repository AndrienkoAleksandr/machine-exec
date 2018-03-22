package websocket

import (
	"bufio"
	"errors"
	"fmt"
	execManager "github.com/AndrienkoAleksandr/machine-exec/exec"
	"github.com/eclipse/che-lib/websocket"
	"github.com/eclipse/che/agents/go-agents/core/rest"
	"log"
	"net"
	"net/http"
	"strconv"
)

const (
	bufferSize int = 8192
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
	id, err := strconv.Atoi(restParmas.Get("id"))
	if err != nil {
		return errors.New("Failed to parse id")
	}
	fmt.Println("Parsed id", id)

	hjr, err := execManager.Attach(id)
	if err != nil {
		return err
	}

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error " + err.Error())
		return err
	}

	hjrConn := hjr.Conn
	hjrReader := hjr.Reader

	defer hjrConn.Close()

	go sendClientInputToExec(hjrConn, wsConn)
	sendExecOutPutToConnection(hjrReader, wsConn)

	return nil
}

func sendClientInputToExec(hjrConn net.Conn, wsConn *websocket.Conn) {
	for {
		msgType, wsBytes, err := wsConn.ReadMessage() //todo check websocket message type

		fmt.Println("userDate='" + string(wsBytes) + "'")

		if err != nil {
			fmt.Println("failed to get read websocket message")
			return
		}

		if msgType != websocket.TextMessage {
			continue
		}

		if hjrConn.Write(wsBytes); err != nil {
			fmt.Println("failed to write client content to exec!!! Cause: " + err.Error())
			return
		}
	}
}

// Todo RuneReader?
func sendExecOutPutToConnection(hjReader *bufio.Reader, wsConn *websocket.Conn) {
	//deferer stop reading exec output! save place to next reading...

	//execBytes := make([]byte, bufferSize)
	//for {
	//	size, err := hjReader.Read(execBytes);
	//	hjReader.ReadByte()
	//	if  err != nil {
	//		fmt.Println("failed to read exec stdOut stream!!! " + err.Error())
	//		return
	//	}
	//
	//	fmt.Println("size=", size)
	//	fmt.Println("exec response='" + string(execBytes[0:size]) +"'")
	//
	//	if size > 0 {
	//		if err := wsConn.WriteMessage(websocket.TextMessage, execBytes[0:size]); err != nil {
	//			fmt.Println("failed to write to websocket message!!!" + err.Error())
	//		}
	//	}
	//}

	for {
		runa, size, err := hjReader.ReadRune()
		if err != nil {
			fmt.Println("failed to read exec stdOut stream!!! " + err.Error())
			return
		}

		fmt.Println("size=", size)
		fmt.Println("exec response='" + string(runa) + "'")

		if size > 0 {
			if err := wsConn.WriteMessage(websocket.TextMessage, []byte(string(runa))); err != nil {
				fmt.Println("failed to write to websocket message!!!" + err.Error())
				return
			}
		}
	}
}
