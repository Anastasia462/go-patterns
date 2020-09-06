package factory_method

// Creature ...
type Creature interface {
	SayHi() string
}

// Human ...
type Human struct {
	words string
}

func (h *Human) SayHi() string {
	return h.words
}

// Cat ...
type Cat struct {
	sound string
}

func (c *Cat) SayHi() string {
	return c.sound
}

// Nastya ...
type Nastya struct {
	words string
}

func (n *Nastya) SayHi() string{
	return n.words
}
