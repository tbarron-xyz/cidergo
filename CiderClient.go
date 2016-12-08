package cidergo

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type CiderClient struct {
	ws *websocket.Conn
}

func NewCiderClient(url string, port int) (client *CiderClient, err error) {
	verbose("Connecting... ")
	ws, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%d/", url, port), nil)
	if err != nil {
		verbose("Error: %s", err.Error())
		return
	}
	verbose("Connected.")
	client = &CiderClient{
		ws,
	}
	return
}

// Sends and returns the raw response
func (c *CiderClient) SendRaw(cmd string) (response []byte, err error) {
	response, err = writeReadOnce(c.ws, cmd)
	return
}

// Sends and gets a Response.
func (c *CiderClient) Send(cmd string) (itf interface{}, err error) {
	var raw []byte
	raw, err = c.SendRaw(cmd)
	if err != nil {
		return
	}
	var response StatusResponse
	err = json.Unmarshal(raw, &response)
	if err != nil {
		return
	}
	itf, err = response.Parse()
	return
}

// Sends, and returns an int
func (c *CiderClient) SendInt(cmd string) (i int, err error) {
	res, err := c.Send(cmd)
	if err != nil {
		return
	}
	i, err = itfToInt(res)
	return
}

// Sends, and returns a string
func (c *CiderClient) SendString(cmd string) (s string, err error) {
	res, err := c.Send(cmd)
	if err != nil {
		return
	}
	s, err = itfToString(res)
	return
}

// Sends, and returns a []string
func (c *CiderClient) SendArrayString(cmd string) (s []string, err error) {
	res, err := c.Send(cmd)
	if err != nil {
		return
	}
	s, err = itfToArrayString(res)
	return
}

// Sends, and returns a bool
func (c *CiderClient) SendBool(cmd string) (b bool, err error) {
	res, err := c.Send(cmd)
	if err != nil {
		return
	}
	b, err = itfToBool(res)
	return
}
