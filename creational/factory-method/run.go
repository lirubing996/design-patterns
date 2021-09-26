package factory_method

import "fmt"


func Run() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	fmt.Println(ak47.toString())
	fmt.Println(musket.toString())
}