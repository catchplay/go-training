package person

type Person struct {
	name string
	age  int
}

func New(name string, age int) *Person {
	return &Person{name, age}
}

func (p Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetAge(age int) {
	p.age = age
}
