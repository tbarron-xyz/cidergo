package cidergo

// Sets the value of key
func (c *CiderClient) Set(key, value string) (err error) {
	_, err = c.Send(SprintfEscape("SET %v %v", key, value))
	return
}

// Gets the value of key
func (c *CiderClient) Get(key string) (value string, err error) {
	value, err = c.SendString(SprintfEscape("GET %v", key))
	return
}

// Appends to the existing value of key
func (c *CiderClient) Append(key, addition string) (value string, err error) {
	value, err = c.SendString(SprintfEscape("APPEND %v %v", key, addition))
	return
}
