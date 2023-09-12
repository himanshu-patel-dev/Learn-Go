package main

// implementation of Shoe
type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

// classes of different shoes

type AdidasShoe struct {
	Shoe
}

type NikeShoe struct {
	Shoe
}
