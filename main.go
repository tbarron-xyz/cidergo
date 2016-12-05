package cidergo

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type StatusResponse struct {
	Status   string
	Response interface{}
	Error    string
}

func (res *StatusResponse) Parse() (itf interface{}, err error) {
	if res.Status == "success" {
		itf = res.Response
	} else {
		err = fmt.Errorf(res.Error)
	}
	return
}

var Txtm = websocket.TextMessage
