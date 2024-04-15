package bar

import "fmt"

type Bar struct{}

func (b *Bar) SayHello()  { fmt.Println("Hello!") }
func (b *Bar) SayHolla()  { fmt.Println("Holla!") }
func (b *Bar) SayPrivet() { fmt.Println("Privet!") }
func (b *Bar) SayGoddag() { fmt.Println("God dag!") }
