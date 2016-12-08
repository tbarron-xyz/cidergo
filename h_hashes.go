package cidergo

// Gets the value at key in hash
func (c *CiderClient) Hget(hash, key string) (value string, err error) {
	value, err = c.SendString(SprintfEscape("HGET %v %v", hash, key))
	return
}

// Sets the value at key in hash
func (c *CiderClient) Hset(hash, key, value string) (err error) {
	_, err = c.Send(SprintfEscape("HSET %v %v %v", hash, key, value))
	return
}

// The number of keys in hash
func (c *CiderClient) Hlen(hash string) (length int, err error) {
	length, err = c.SendInt(SprintfEscape("HLEN %v %v %v", hash))
	return
}

// The keys in hash
func (c *CiderClient) Hkeys(hash string) (keys []string, err error) {
	keys, err = c.SendArrayString(SprintfEscape("HKEYS %v", hash))
	return
}
