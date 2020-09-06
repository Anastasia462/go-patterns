package factory_method

// Creator ...
type Creator interface {
	CreateCreature(creature string) Creature
}

// ConcreteCreator ...
type ConcreteCreator struct {}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (c *ConcreteCreator) CreateCreature(newCreature string) Creature {
	var creature Creature

	switch newCreature {
	case "human":
		creature = &Human{"hi"}
	case "cat":
		creature = &Cat{"meow"}
	case "nastya":
		creature = &Nastya{"i'm stupid"}

	}

	return creature
}

