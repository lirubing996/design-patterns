package abstract_factory

import "fmt"

type iShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
	toString() string
}

type shirt struct {
	logo string
	size int
}

type adidasShirt struct {
	shirt
}

type nikeShirt struct {
	shirt
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) getSize() int {
	return s.size
}

func (s *shirt) toString() string {
	return fmt.Sprintf("The %s's shirt size is %d", s.logo, s.size)
}