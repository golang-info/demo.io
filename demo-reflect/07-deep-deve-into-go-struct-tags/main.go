package main

/*
	Deep dive into Go Struct Tags
	https://medium.com/@metamemelord/deep-dive-into-go-struct-tags-d629ec0be30d
*/

import (
	"fmt"
	"reflect"
	"time"
)

type Todo struct {
	ID  int `db:"id"`
	Content string `db:"content"`
	Timestamp   int64 `db:"ts"`
}

func NewTodo(id int, content string, timestamp int64) *Todo {
	return &Todo{
		ID: id,
		Content: content,
		Timestamp: timestamp,
	}
}

func main() {
	todo := NewTodo(1, "Chop carroots", time.Now().Unix())

	v := reflect.ValueOf(todo)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("Not a struct!")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(
			t.Field(i).Name,
			t.Field(i).Tag.Get("db"),
			v.Field(i).Interface())
	}


}
