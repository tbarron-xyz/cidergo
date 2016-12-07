package cidergo

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

var Verbose = false

func verbose(args ...interface{}) {
	if Verbose {
		fmt.Println(args...)
	}
}

func SprintfEscape(format string, args ...string) (formatted string) {
	escapedArgs := []interface{}{}
	for _, e := range args {
		escapedArgs = append(escapedArgs, escape(e))
	}
	formatted = fmt.Sprintf(format, escapedArgs...)
	return
}

func readOnce(ws *websocket.Conn) (msg []byte, err error) {
	_, msg, err = ws.ReadMessage()
	if err != nil {
		verbose("Read error: " + err.Error())
		ws.Close()
	}
	// verbose("<", string(msg))
	return
}

func writeOnce(ws *websocket.Conn, cmd string) (err error) {
	err = ws.WriteMessage(txtm, []byte(strings.TrimSpace(cmd)))
	if err != nil {
		verbose("Write error: " + err.Error())
		ws.Close()
	}
	return
}

func writeReadOnce(ws *websocket.Conn, cmd string) (msg []byte, err error) {
	err = writeOnce(ws, cmd)
	if err != nil {
		return
	}
	msg, err = readOnce(ws)
	return
}

func escape(input string) (output string) {
	var buffer = &bytes.Buffer{}
	buffer.WriteRune('"')
	for _, e := range input {
		if e == '"' {
			buffer.WriteRune('\\')
		}
		buffer.WriteRune(e)
	}
	buffer.WriteRune('"')
	output = buffer.String()
	return
}

func itfToString(itf interface{}) (s string, err error) {
	switch itf.(type) {
	case string:
		s = itf.(string)
	default:
		err = fmt.Errorf("Not a string.")
	}
	return
}

func itfToInt(itf interface{}) (i int, err error) {
	switch itf.(type) {
	case int:
		i = itf.(int)
	default:
		err = fmt.Errorf("Not an int.")
	}
	return
}
