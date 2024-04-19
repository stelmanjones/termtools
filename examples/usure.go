package examples

import "github.com/stelmanjones/termtools/usure"

type Dog struct {
	name string
	age  int
}

func usureExample() {
	first := Dog{"Fido", 1}
	second := Dog{"Daisy", 3}
	usure.ExpectEqual("dogs are not equal", first, second)

}
