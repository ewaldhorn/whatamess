package main

type Employee struct {
	Name      string
	Role      string
	MinSalary int
	MaxSalary int
}

type Builder struct {
	e Employee
}

func (b *Builder) Build() Employee {
	return b.e
}

func (b *Builder) Name(name string) *Builder {
	b.e.Name = name
	return b
}

func (b *Builder) Role(role string) *Builder {
	if role == "manager" {
		b.e.MinSalary = 20000
		b.e.MaxSalary = 60000
	}
	b.e.Role = role
	return b
}
