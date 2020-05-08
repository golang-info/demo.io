package main

import "testing"

/**
	https://blog.csdn.net/luopotaotao/article/details/79192657
	装饰类与要装饰的目标实现相同的接口， 并持有原始对象的引用
	依据李氏替换原则， 装饰器可以替换原始对象
	而在装饰器类的接口实现中，可以添加额外的逻辑， 在进行原始对象的调用或者不调用，从而实现对原始对象所实现的接口的装饰增强
	在Java中，AOP本质上就是通过这种方式拦截函数调用， 对其进行装饰后调用装饰之后的方法
*/

func TestDecorate(t *testing.T) {
	//最原始的功能， 实现了Show接口
	xiaocai := &Person{"小菜"}

	//用以对原始功能进行增强的新对象
	finery := &Finery{}
	tshirts := &TShirts{}
	sneakers := &Sneakers{}

	//对原始功能进行了三次装饰增强
	finery.Decorate(xiaocai)
	tshirts.Decorate(finery)
	sneakers.Decorate(tshirts)

	//调用装饰增强之后的实现
	sneakers.Show()

}
