package websocket

import (
	"bufio"
	"bytes"
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
	bufferSize = 8192
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
		return errors.New("failed to parse id")
	}
	fmt.Println("Parsed id", id)

	hjr, err := execManager.Attach(id)
	if err != nil {
		return err
	}

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Unable to upgrade connection to websocket " + err.Error())
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
		msgType, wsBytes, err := wsConn.ReadMessage()
		if err != nil {
			fmt.Println("failed to get read websocket message")
			return
		}

		if msgType != websocket.TextMessage {
			continue
		}

		if hjrConn.Write(wsBytes); err != nil {
			fmt.Println("failed to write client content to the exec!!! Cause: " + err.Error())
			return
		}
	}
}

func sendExecOutPutToConnection(hjReader *bufio.Reader, wsConn *websocket.Conn) {
	buf := make([]byte, bufferSize)
	var buffer bytes.Buffer

	for {
		rbBize, err := hjReader.Read(buf)
		if err != nil {
			fmt.Println("failed to read exec stdOut/stdError stream!!! " + err.Error())
			return
		}

		i, err := normalizeBuffer(&buffer, buf, rbBize)
		if err != nil {
			log.Printf("Couldn't normalize byte buffer to UTF-8 sequence, due to an error: %s", err.Error())
			return
		}

		if rbBize > 0 {
			if err := wsConn.WriteMessage(websocket.TextMessage, buffer.Bytes()); err != nil {
				fmt.Println("failed to write to websocket message!!!" + err.Error())
				return
			}
		}

		buffer.Reset()
		if i < rbBize {
			buffer.Write(buf[i:rbBize])
		}
	}
}
