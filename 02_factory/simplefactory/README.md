# 简单工厂模式
**简单工厂模式（Simple Factory Pattern）:** 定义一个工厂类，它可以根据参数的不同返回不同类的实例，被创建的实例通常都具有共同的父类。因为在简单工厂模式中用于创建实例的方法是静态（static）方法，因此简单工厂模式又被称为静态工厂方法（Static Factory Method）模式，它属于类创建型模式。

简单工厂的要点在于：当你需要什么对象，只需要传入一个正确的参数，就可以获取你所需要的对象，而无需知道其创建细节，其核心是工厂类的设计，其结构如图所示：

例如：
我们创建一个工厂结构体，并且创建一个产品接口，工厂可以创建产品，只要在工厂的某个方法中传入不同的参数，就可以返回实现产品接口的不同的对象

1.创建工厂结构体:
```go
type Factory struct {}
```
2.创建产品接口，这里为了方便，只写了一个方法，请根据自己的需要扩展
```go
type Product interface {
    create()
}
```
3.创建两个产品：产品1和产品2，它们实现了产品接口:
```go
type Product1 struct{}

func (p *Product1) create() string {
	return "this is product1"
}

type Product2 struct{}

func (p *Product2) create() string {
	return "this is product2"
}
```
4.为工厂结构体添加一个方法用于生产产品（实例化对象）:
```go
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
```

这样就可以通过传入不同的方法得到不同的产品实例了:
```go
    // 创建一个工厂类，在应用中可以将这个工厂类实例作为一个全局变量
    factory := new(Factory)
    
    // 在工厂类中传入不同的参数，获取不同的实例
    p1 := factory.Generate("product1")
    p1.create()                          // output:   this is product 1
    
    p2 := factory.Generate("product2")
    p2.create()                          // output:   this is product 2
```

下面说说工厂模式的优缺点:

**优点:** 工厂类是整个工厂模式的核心，我们只需要传入给定的信息，就可以创建所需实例，在多人协作的时候，无需知道对象之间的内部依赖，可以直接创建，有利于整个软件体系结构的优化

**缺点:** 工厂类中包含了所有实例的创建逻辑，一旦这个工厂类出现问题，所有实例都会受到影响，并且，工厂类中生产的产品都基于一个共同的接口，一旦要添加不同种类的产品，这就会增加工厂类的复杂度，将不同种类的产品混合在一起，违背了单一职责，系统的灵活性和可维护性都会降低，并且当新增产品的时候，必须要修改工厂类，违背了『系统对扩展开放，对修改关闭』的原则