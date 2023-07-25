package puppy

import (
	"fmt"

	"github.com/vietdv277/dog"
)

func Bark() string {
	return "Wolf!"
}

func Barks() string {
	return "Wolf.....Wolf.....Wolf"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("Message from version 1.1.0")
}
