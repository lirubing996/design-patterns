package abstract_factory

import "fmt"

type iShoe interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
	toString() string
}

type shoe struct {
	logo string
	size int
}

type adidasShoe struct {
	shoe
}

type nikeShoe struct {
	shoe
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) getSize() int {
	return s.size
}

func (s *shoe) toString() string {
	return fmt.Sprintf("The %s's shoe size is %d", s.logo, s.size)
}