package main

type human interface {
	speak() string
}

type man struct {
	name string
	age int
}

func (m man) speak() string {
	//k := fmt.Sprintf("Hi, i am %v, age %v", m.name, m.age)
	return m.name
}

func dosomething(h human) {
	h.speak()
}

func main() {
	a := man{"xiaoming", 11}
	
	dosomething(a)
}