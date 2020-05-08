package simple_factory

import "fmt"

// API is interface
type API interface {
	Say(name string) string
}

//NewAPI return Api instance my type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else {
		return &helloAPI{}
	}
	return nil
}

//hiAPI is one of API implement
type hiAPI struct {

}

//Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//HelloAPI is another API implement
type helloAPI struct {

}

//Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
