package abstract_factory

type NikeFactory struct {
}

func NewNikeFactory() AbstractFactory {
	return &NikeFactory{}
}

func (f *NikeFactory) CreateShirt(material string, size int) AbstractShirt {
	return &NikeShirt{
		material: material,
		size:     size,
	}
}

func (f *NikeFactory) CreateShorts(material string, size int) AbstractShorts {
	return &NikeShorts{
		material: material,
		size:     size,
	}
}

type NikeShirt struct {
	material string
	size     int
}

func (s *NikeShirt) SetMaterial(material string) {
	s.material = material
}

func (s *NikeShirt) GetMaterial() (material string) {
	return s.material
}

func (s *NikeShirt) SetSize(size int) {
	s.size = size
}

func (s *NikeShirt) GetSize() (size int) {
	return s.size
}

type NikeShorts struct {
	material string
	size     int
}

func (s *NikeShorts) SetMaterial(material string) {
	s.material = material
}

func (s *NikeShorts) GetMaterial() (material string) {
	return s.material
}

func (s *NikeShorts) SetSize(size int) {
	s.size = size
}

func (s *NikeShorts) GetSize() (size int) {
	return s.size
}
