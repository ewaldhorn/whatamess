package foo

type BarInterface interface{ SayHello() }
type Foo struct{ bar BarInterface }

func NewFoo(bar BarInterface) *Foo { return &Foo{bar: bar} }
func (f *Foo) Greet()              { f.bar.SayHello() }
