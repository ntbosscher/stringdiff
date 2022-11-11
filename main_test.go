package stringdiff

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	df := New("Hello world, this is a string", "Hello world, we have something new")
	fmt.Println(df)
	// > Hello world,
	// - this is a string
	// + we have something new

	df = New("Hello John, this is a string", "Hello Kennith Watson, this is a string")
	fmt.Println(df)
	// > Hello
	// - John,
	// + Kennith Watson,
	// > this is a string

	df = New("Hello John, how are you?", "Hello Kennith Watson, how is your fiance?")
	fmt.Println(df)
	// > Hello
	// - John,
	// + Kennith Watson,
	// > how
	// - are you?
	// + is your fiance
}
