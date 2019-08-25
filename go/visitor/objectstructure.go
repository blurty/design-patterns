package visitor

import (
	"container/list"
)

func NewObject() *object {
	return &object{
		list: list.New(),
	}
}

type object struct {
	list *list.List
}

func (o *object) Attach(a Acceptor) *list.Element {
	return o.list.PushBack(a)
}

func (o *object) Detach(e *list.Element) {
	o.list.Remove(e)
}

func (o *object) Accept(visitor Visitor) {
	for e := o.list.Front(); e != nil; e = e.Next() {
		if acc, ok := e.Value.(Acceptor); ok {
			acc.Accept(visitor)
		}
	}
}
