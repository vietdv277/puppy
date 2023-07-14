package puppy

import "github.com/vietdv277/dog"

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
