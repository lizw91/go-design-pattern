package simplefactory

// 创建工厂结构体
type Factory struct{}

// 创建产品接口，这里为了方便只加了一个方法
type Product interface {
	create() string
}

// 创建两个产品：产品1和产品2，它们实现了产品接口
type Product1 struct{}

func (p *Product1) create() string {
	return "this is product1"
}

type Product2 struct{}

func (p *Product2) create() string {
	return "this is product2"
}

// 为工厂结构体添加一个方法用于生产产品（实例化对象）:
func (f *Factory) Generate(name string) Product {
	switch name {
	case "product1":
		return &Product1{}
	case "product2":
		return &Product2{}
	default:
		return nil
	}
}

func Run(name string) string {
	// 创建一个工厂类，在应用中可以将这个工厂类实例作为一个全局变量
	f := Factory{}

	// 在工厂类中传入不同的参数，获取不同的实例
	p := f.Generate(name)
	return p.create()
}
