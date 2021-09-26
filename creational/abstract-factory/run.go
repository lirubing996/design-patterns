package abstract_factory

import (
    "fmt"
    "github.com/lirubing996/design-patterns/tools"
)

func Run() {
    adidasFactory, _ := getSportsFactory("adidas")
    nikeFactory, _ := getSportsFactory("nike")

    nikeShoe := nikeFactory.makeShoe()
    nikeShirt := nikeFactory.makeShirt()
    fmt.Println(nikeShoe.toString())
    fmt.Println(nikeShirt.toString())

    tools.PrintHr()

    adidasShoe := adidasFactory.makeShoe()
    adidasShirt := adidasFactory.makeShirt()
    fmt.Println(adidasShoe.toString())
    fmt.Println(adidasShirt.toString())
}
