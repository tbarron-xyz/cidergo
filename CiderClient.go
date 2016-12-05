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
	var response Response
	err = json.Unmarshal(raw, &response)
	if err != nil {
		return
	}
	itf, err = response.Parse()
	return
}

func (c *CiderClient) SendInt(cmd string) (i int, err error) {
	res, err := c.Send(cmd)
	if err != nil {
		return
	}
	i, err = itfToInt(res)
	return
}

func (c *CiderClient) SendString(cmd string) (s string, err error) {
	res, err := c.Send(cmd)
	if err != nil {
		return
	}
	s, err = itfToString(res)
	return
}

func (c *CiderClient) Set(key, value string) (err error) {
	_, err = c.Send(SprintfEscape("SET %v %v", key, value))
	return
}

func (c *CiderClient) Get(key string) (value string, err error) {
	value, err = c.SendString(SprintfEscape("GET %v", key))
	return
}

func (c *CiderClient) Hget(hash, key string) (value string, err error) {
	value, err = c.SendString(SprintfEscape("HGET %v %v", hash, key))
	return
}

func (c *CiderClient) Hset(hash, key, value string) (err error) {
	_, err = c.Send(SprintfEscape("HSET %v %v %v", hash, key, value))
	return
}
