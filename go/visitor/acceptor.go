package visitor

type Acceptor interface {
	Accept(visitor Visitor)
	GetName() string
}

func NewMan() *man {
	return &man{
		name: "男人",
	}
}

type man struct {
	name string
}

func (m *man) Accept(visit Visitor) {
	visit.VisitMan(m)
}

func (m *man) GetName() string {
	return m.name
}

func NewWoman() *woman {
	return &woman{
		name: "女人",
	}
}

type woman struct {
	name string
}

func (w *woman) Accept(visit Visitor) {
	visit.VisitWoman(w)
}

func (w *woman) GetName() string {
	return w.name
}
