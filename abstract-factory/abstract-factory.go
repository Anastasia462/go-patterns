package abstract_factory

// AbstractFactory ...
type AbstractFactory interface {
	CreateShirt(material string, size int) AbstractShirt
	CreateShorts(material string, size int) AbstractShorts
}

// AbstractShirt ...
type AbstractShirt interface {
	SetMaterial(material string)
	GetMaterial() string
	SetSize(size int)
	GetSize() int
}

// AbstractShorts ...
type AbstractShorts interface {
	SetMaterial(material string)
	GetMaterial() string
	SetSize(size int)
	GetSize() int
}
