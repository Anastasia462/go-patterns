package abstract_factory

type AdidasFactory struct {
}

func NewAdidasFactory() AbstractFactory {
	return &AdidasFactory{}
}

func (f *AdidasFactory) CreateShirt(material string, size int) AbstractShirt {
	return &AdidasShirt{
		material: material,
		size:     size,
	}
}

func (f *AdidasFactory) CreateShorts(material string, size int) AbstractShorts {
	return &AdidasShorts{
		material: material,
		size:     size,
	}
}

type AdidasShirt struct {
	material string
	size     int
}

func (s *AdidasShirt) SetMaterial(material string) {
	s.material = "chinese" + material
}

func (s *AdidasShirt) GetMaterial() string {
	return "chinese" + s.material
}

func (s *AdidasShirt) SetSize(size int) {
	s.size = size + 2
}

func (s *AdidasShirt) GetSize() int {
	return s.size + 2
}

type AdidasShorts struct {
	material string
	size     int
}

func (s *AdidasShorts) SetMaterial(material string) {
	s.material = "chinese" + material
}

func (s *AdidasShorts) GetMaterial() string {
	return "chinese" + s.material
}

func (s *AdidasShorts) SetSize(size int) {
	s.size = size + 2
}

func (s *AdidasShorts) GetSize() int {
	return s.size + 2
}
