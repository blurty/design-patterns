package main

import "github.com/blurty/design-pattens/go/visitor"

func main() {
	o := visitor.NewObject()
	m := visitor.NewMan()
	w := visitor.NewWoman()
	mItem := o.Attach(m)
	wItem := o.Attach(w)

	v := visitor.NewSuccess()
	o.Accept(v)

	// 解除
	o.Detach(mItem)
	o.Detach(wItem)
	o.Accept(v)
}
