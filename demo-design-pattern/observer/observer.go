package observer

import "fmt"

// 很多东西不是会了才能做， 而是做了才能学会

// 报社 -- 客户
type Customer interface {
	update()
}

type CustomerA struct {

}

func (c *CustomerA) update() {
	fmt.Println("我是客户A， 我收到报纸了")
}

type CustomerB struct {

}

func (c *CustomerB) update() {
	fmt.Println("我是客户B， 我收到报纸了")
}

// 报社 （被观察者）
type NewOffice struct {
	customers []Customer
}

func (n *NewOffice) addCustomer(customer Customer) {
	n.customers = append(n.customers, customer)
}

func (n *NewOffice) newspaperCome() {
	// 通知所有客户
	n.notifyAllCustomer()
}

func (n *NewOffice) notifyAllCustomer() {
	for _, customer := range n.customers {
		customer.update()
	}
}