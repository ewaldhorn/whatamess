type Person struct {
	Name string
	Age  int
}

func WithName(name string) func(*Person) {
	return func(p *Person) {
		p.Name = name
	}
}

func WithAge(age int) func(*Person) {
	return func(p *Person) {
		p.Age = age
	}
}

func main() {
	person := Build(Person{}, WithName("John"), WithAge(30))
	fmt.Println(person) // Output: {John 30}
}
