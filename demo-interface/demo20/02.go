package main

import "fmt"

type Stringer interface {
	String() string
}

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	// fmt.Println(a.String())
	Print(a)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
