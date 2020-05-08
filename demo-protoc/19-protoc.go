package main

import (
	"fmt"
	"log"

	"demo.io/demo-protoc/test"
	"github.com/golang/protobuf/proto"
)

func main() {
	student := &test.Student{
		Name:   "test",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}
	data, err := proto.Marshal(student)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println("data: ", data)

	newStudent := &test.Student{}
	err = proto.Unmarshal(data, newStudent)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(student.GetName(), newStudent.GetName())

	//Now test and newTest contain the same data
	if student.GetName() != newStudent.GetName() {
		log.Fatal("data mismatch %q != %q", student.GetName(), newStudent.GetName())
	}
}
