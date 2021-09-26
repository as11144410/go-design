package go_design

type Product interface {
	SetName(name string)
	GetName() string
}

// Product1 产品1
type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name = name
}

func (p1 *Product1) GetName() string {
	return "产品1的name为:" + p1.name
}

// Product2 产品2
type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name = name
}

func (p2 *Product2) GetName() string {
	return "产品2的name为:" + p2.name
}

// 实现简单工厂类

type productType int

const (
	p1 productType = iota
	p2
)

type productFactory struct {
}

func (pf *productFactory) Create(productType productType) Product {
	if productType == p1 {
		return &Product1{}
	}
	if productType == p2 {
		return &Product2{}
	}
	return nil
}
