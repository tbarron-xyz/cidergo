package cidergo

// Add an element
func (c *CiderClient) Sadd(set, element string) (added int, err error) {
	added, err = c.SendInt(SprintfEscape("SADD %v %v", set, element))
	return
}

// Remove an element
func (c *CiderClient) Srem(set, element string) (removed int, err error) {
	removed, err = c.SendInt(SprintfEscape("SREM %v %v", set, element))
	return
}

// Is element a member of set?
func (c *CiderClient) Sismember(set, element string) (isElement bool, err error) {
	isElement, err = c.SendBool(SprintfEscape("SISMEMBER %v %v", set, element))
	return
}

// Cardinality
func (c *CiderClient) Scard(set string) (card int, err error) {
	card, err = c.SendInt(SprintfEscape("SISMEMBER %v", set))
	return
}

// Pops a random element
func (c *CiderClient) Spop(set string) (element string, err error) {
	element, err = c.SendString(SprintfEscape("SPOP %v", set))
	return
}

// Returns (and does not remove) a random element
func (c *CiderClient) Srandmember(set string) (element string, err error) {
	element, err = c.SendString(SprintfEscape("SRANDMEMBER %v", set))
	return
}
