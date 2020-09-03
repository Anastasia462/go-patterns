package builder

// Builder ...
type Builder interface {
	CreateHead(str string)
	CreateBody(str string)
	CreateFoot(str string)
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.CreateHead("smart")
	d.builder.CreateBody("flawless")
	d.builder.CreateFoot("strong")
}
