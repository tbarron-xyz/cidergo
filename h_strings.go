package cidergo

func (c *CiderClient) Set(key, value string) (err error) {
	_, err = c.Send(SprintfEscape("SET %v %v", key, value))
	return
}

func (c *CiderClient) Get(key string) (value string, err error) {
	value, err = c.SendString(SprintfEscape("GET %v", key))
	return
}

func (c *CiderClient) Append(key, addition string) (value string, err error) {
	value, err = c.SendString(SprintfEscape("APPEND %v %v", key, addition))
	return
}
