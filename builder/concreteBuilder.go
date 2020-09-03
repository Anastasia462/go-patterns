package builder

type ConcreteBuilder struct {
	human *Human
}

func (b *ConcreteBuilder) CreateHead(headType string) {
	b.human.Parts += headType + " head\n"
}

func (b *ConcreteBuilder) CreateBody(bodyType string) {
	b.human.Parts += bodyType + " body\n"
}

func (b *ConcreteBuilder) CreateFoot(footType string) {
	b.human.Parts += footType + " foot\n"
}
