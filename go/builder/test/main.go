package main

import "github.com/blurty/design-pattens/go/builder"

func main() {
	b := builder.NewThinPersonBuilder()
	builder.DirectBuildPerson(b)
	person := b.GetPerson()
	person.Show()
}
