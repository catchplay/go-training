package person

type Person struct {
	name string
	age  int
}

func New(name string, age int) *Person {
	return &Person{name, age}
}

func (p *Person) Name() string {
	return p.name
}

func (p Person) SetName(name string) {
	p.name = name
}

func (p *Person) Age() int {
	return p.age
}
func (p *Person) SetAge(age int) {
	p.age = age
}
