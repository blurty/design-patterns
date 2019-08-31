package builder

import "fmt"

type Person struct {
	Parts []string
}

func (p *Person) Add(part string) {
	p.Parts = append(p.Parts, part)
}

func (p *Person) Show() {
	for _, part := range p.Parts {
		fmt.Println(part)
	}
}

type PersonBuider interface {
	BuildHead()
	BuildBody()
	BuildArms()
	BuildLegs()
}

func NewThinPersonBuilder() *ThinPersonBuilder {
	return &ThinPersonBuilder{}
}

type ThinPersonBuilder struct {
	p Person
}

func (b *ThinPersonBuilder) BuildHead() {
	b.p.Add("draw head ...")
	return
}

func (b *ThinPersonBuilder) BuildBody() {
	b.p.Add("draw body ...")
	return
}

func (b *ThinPersonBuilder) BuildArms() {
	b.p.Add("draw arms ...")
	return
}

func (b *ThinPersonBuilder) BuildLegs() {
	b.p.Add("draw legs ...")
	return
}

func (b *ThinPersonBuilder) GetPerson() *Person {
	return &b.p
}
