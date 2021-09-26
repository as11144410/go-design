package go_design

import (
	"fmt"
	"testing"
)

// 直接创建
func TestProduct_Create(t *testing.T) {
	product1 := &Product1{}
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := &Product2{}
	product2.SetName("p2")
	fmt.Println(product2.GetName())
}

// 工厂创建
func TestProductFactory_Create(t *testing.T) {
	productFactory := productFactory{}
	product1 := productFactory.Create(p1)
	product1.SetName("p1")
	fmt.Println(product1.GetName())
	product2 := productFactory.Create(p2)
	product2.SetName("p2")
	fmt.Println(product2.GetName())

}
