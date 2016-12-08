package cidergo

import "strconv"

// Increments key by incr, and decrements it after duration
func (c *CiderClient) Incrby(key string, incr int) (value int, err error) {
	value, err = c.SendInt(SprintfEscape("INCRBY %v %v", key, strconv.Itoa(incr)))
	return
}

// Sets the duration that increments on key last before being decremented
func (c *CiderClient) Duration(key string, duration int) (err error) {
	_, err = c.Send(SprintfEscape("DURATION %v %v", key, strconv.Itoa(duration)))
	return
}

// Resets the value of key to zero (does not change existing duration)
func (c *CiderClient) Creset(key string) (err error) {
	_, err = c.Send(SprintfEscape("CRESET %v", key))
	return
}

// Gets the value of key
func (c *CiderClient) Cget(key string) (value int, err error) {
	value, err = c.SendInt(SprintfEscape("CGET %v", key))
	return
}
