package cidergo

func (c *CiderClient) Lpush(key, element string) (length int, err error) {
	length, err = c.SendInt(SprintfEscape("LPUSH %v %v", key, element))
	return
}

func (c *CiderClient) Rpush(key, element string) (length int, err error) {
	length, err = c.SendInt(SprintfEscape("RPUSH %v %v", key, element))
	return
}

func (c *CiderClient) Lpop(key string) (element string, err error) {
	element, err = c.SendString(SprintfEscape("LPOP %v", key))
	return
}

func (c *CiderClient) Rpop(key string) (element string, err error) {
	element, err = c.SendString(SprintfEscape("RPOP %v", key))
	return
}

func (c *CiderClient) Llen(key string) (length int, err error) {
	length, err = c.SendInt(SprintfEscape("LLEN %v", key))
	return
}
