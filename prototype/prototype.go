package prototype

// Prototyper ...
type Prototyper interface {
	Clone() Prototyper
	GetSize() int
	GetColor() string
}

func NewShoes(size int, color string) Prototyper {
	return &Shoes{
		size:  size,
		color: color,
	}
}

type Shoes struct {
	size int
	color string
}

func (s *Shoes) GetSize() int{
	return s.size
}

func (s *Shoes) GetColor() string {
	return s.color
}

func (s *Shoes) Clone() Prototyper{
	return &Shoes{
		size:  s.size,
		color: s.color,
	}
}
