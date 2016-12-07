package cidergo

import "strconv"

func (c *CiderClient) Incrby(key string, incr int) (value int, err error) {
	value, err = c.SendInt(SprintfEscape("INCRBY %v %v", key, strconv.Itoa(incr)))
	return
}

func (c *CiderClient) Duration(key string, duration int) (err error) {
	_, err = c.Send(SprintfEscape("DURATION %v %v", key, strconv.Itoa(duration)))
	return
}

func (c *CiderClient) Creset(key string) (err error) {
	_, err = c.Send(SprintfEscape("CRESET %v", key))
	return
}

func (c *CiderClient) Cget(key string) (value int, err error) {
	value, err = c.SendInt(SprintfEscape("CGET %v", key))
	return
}
