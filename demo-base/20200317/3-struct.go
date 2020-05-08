package main

import (
	"fmt"
)

type Member struct {
	id int
	name string
	email string
	gender int
	age int
}

type Student struct {
	Member
	class int
}



func Change(m1 Member, m2 *Member) {
	m1.name = "小李"
	m2.name = "小李"
}

func main() {
	m1 :=  Member{
		name: "小王",
	}
	m2 := &Member{
		name: "小王",
	}

	m3 := new(Member)
	m3.name = "小王"

	Change(m1, m2)
	fmt.Println(m1, m2, m3)

	m4_0 := Member{
		id:     1,
		name:   "hhh",
		email:  "hhh@gmail.com",
		gender: 1,
		age:    12,
	}
	m4 := &Student{
		m4_0, 3,
	}
	fmt.Println(m4.name, m4.age, m4.class)



}
