package puppy

import (
	"github.com/skiredall/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUP(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUP(Barks())
}
