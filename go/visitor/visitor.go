package visitor

import "fmt"

type Visitor interface {
	VisitMan(Acceptor)
	VisitWoman(Acceptor)
}

func NewSuccess() *Success {
	return &Success{
		Name: "成功",
	}
}

type Success struct {
	Name string
}

func (s *Success) VisitMan(acceptor Acceptor) {
	fmt.Printf("%s %s时，背后通常有个伟大的女人！\n", acceptor.GetName(), s.Name)
}

func (s *Success) VisitWoman(acceptor Acceptor) {
	fmt.Printf("%s %s时，背后通常有个失败的男人！\n", acceptor.GetName(), s.Name)
}
