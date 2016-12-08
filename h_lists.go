package cidergo

// Push element to the left side of list
func (c *CiderClient) Lpush(key, element string) (length int, err error) {
	length, err = c.SendInt(SprintfEscape("LPUSH %v %v", key, element))
	return
}

// Push element to the right side of list
func (c *CiderClient) Rpush(key, element string) (length int, err error) {
	length, err = c.SendInt(SprintfEscape("RPUSH %v %v", key, element))
	return
}

// Pop element from the left side of list
func (c *CiderClient) Lpop(key string) (element string, err error) {
	element, err = c.SendString(SprintfEscape("LPOP %v", key))
	return
}

// Pop element from the right side of list
func (c *CiderClient) Rpop(key string) (element string, err error) {
	element, err = c.SendString(SprintfEscape("RPOP %v", key))
	return
}

// Length of the list
func (c *CiderClient) Llen(key string) (length int, err error) {
	length, err = c.SendInt(SprintfEscape("LLEN %v", key))
	return
}
