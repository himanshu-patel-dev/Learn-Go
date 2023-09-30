package main

// interface of dress
type Dress interface {
	getColor() string
}

// terrorist dress
type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

// police dress
type PoliceDress struct {
	color string
}

func (c *PoliceDress) getColor() string {
	return c.color
}

func newPoliceDress() *PoliceDress {
	return &PoliceDress{color: "green"}
}
