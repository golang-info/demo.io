package observer

import "testing"

func TestObserver(t *testing.T) {
	CustomerA := &CustomerA{}
	CustomerB := &CustomerB{}

	office := &NewOffice{}

	// 模拟客户订阅
	office.addCustomer(CustomerA)
	office.addCustomer(CustomerB)

	// 新的报纸
	office.newspaperCome()
}
