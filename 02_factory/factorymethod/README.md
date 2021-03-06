# 工厂方法模式

**工厂方法模式（英语：Factory method pattern）:** 是一种实现了“工厂”概念的面向对象设计模式。就像其他创建型模式一样，它也是处理在不指定对象具体类型的情况下创建对象的问题。工厂方法模式的实质是“定义一个创建对象的接口，但让实现这个接口的类来决定实例化哪个类。工厂方法让类的实例化推迟到子类中进行。”

简单工厂需要:
* 1.工厂结构体
* 2.产品接口
* 3.产品结构体

工厂方法需要:
* 1.工厂接口
* 2.工厂结构体
* 3.产品接口
* 4.产品结构体

## 简单工厂与工厂方法区别
在 **简单工厂**中，依赖于唯一的工厂对象，如果我们需要实例化一个产品，那么就要向工厂中传入一个参数获取对应对象，如果要增加一种产品，就要在工厂中修改创建产品的函数，耦合性过高
，而在 **工厂方法** 中，依赖工厂接口，我们可以通过实现工厂接口，创建多种工厂，将对象创建由一个对象负责所有具体类的实例化，变成由一群子类来负责对具体类的实例化，将过程解耦。

## 例子
我现在想在我的老家齐齐哈尔开一家包子店，卖猪肉馅和三鲜馅两种馅料的包子，那么我们使用简单工厂模式应该怎样实现呢？
### 简单工厂实现
#### 1.创建工厂结构体（包子店）
```go
 // 工厂类(包子店) 
   type BunShop struct {}
```
#### 2.创建产品接口（包子类的接口）
```go
  type Bun interface {
      create()
  }
```
#### 3.实现产品（2种包子）
```go
  type PigMeatBuns struct {}

  func (p PigMeatBuns) create() {
    fmt.Println("猪肉馅包子")
  }

  type SamSunStuffingBuns struct {}

  func (s SamSunStuffingBuns) create() {
    fmt.Println("三鲜馅包子")
  }
```
#### 4.为工厂添加生产包子的方法
```go
  func (b BunShop) Generate(t string) Bun {
    switch t {
      case "pig":
        return PigMeatBuns{}
      case "3s":
        return SamSunStuffingBuns{}
      default:
        return nil
    }
  }
```
#### 5.这样一个简单工厂就完成了
```go
  factory := new(BunShop)
  b := factory.Generate("pig")
  b.create()                   // output: 猪肉馅包子
```
可是如果生意做的不错，我想要在广东开一家分店该怎么办呢？依旧是两种包子，但是为了符合当地人的口味，一定会有所差别，难道要一步一步的修改工厂类吗？

这样工厂方法模式就派上用场了...

### 工厂方法实现
#### 1.添加工厂接口(包子店的接口)和产品接口(包子接口)
```go
  type BunShopInterface interface{
     Generate(t string) Bun
  }

  type Bun interface {
     create()
  }
```
#### 2.创建工厂结构体和产品结构体（具体包子店和具体包子）
```go
// 产品结构体
type QSPigMeatBuns struct{}
type QSSamSunStuffingBuns struct{}

type GDPigMeatBuns struct{}
type GDSamSunStuffingBuns struct{}

// 产品接口的实现
func (q *QSPigMeatBuns) create() string {
	return "this is QSPigMeatBuns"
}

func (q *QSSamSunStuffingBuns) create() string {
	return "this is QSSamSunStuffingBuns"
}

func (g *GDPigMeatBuns) create() string {
	return "this is GDPigMeatBuns"
}

func (g *GDSamSunStuffingBuns) create() string {
	return "this is GDSamSunStuffingBuns"
}
```
#### 3.创建对应的工厂（齐市包子店和广东包子店), 并实现工厂类的接口
```go
// 工厂结构体
type QSBunShop struct{}

type GDBunShop struct{}

// 工厂结构体实现工厂接口
func (q *QSBunShop) Generate(t string) Bun {
	switch t {
	case "pig":
		return &QSPigMeatBuns{}
	case "3s":
		return &QSSamSunStuffingBuns{}
	default:
		return nil
	}
}

func (g *GDBunShop) Generate(t string) Bun {
	switch t {
	case "pig":
		return &GDPigMeatBuns{}
	case "3s":
		return &GDSamSunStuffingBuns{}
	default:
		return nil
	}
}
```
#### 4.这样，就完成了工厂方法模式
```go
  var b Bun
  // 卖呀卖呀卖包子...
  QSFactory := new(QSBunShop)
  b = QSFactory.Generate("pig")  // 传入猪肉馅的参数，会返回齐市包子铺的猪肉馅包子
  b.create()

  GDFactory := new(GDBunShop)
  b = GDFactory.Generate("pig") // 同样传入猪肉馅的参数，会返回广东包子铺的猪肉馅包子
  b.create()
```

## 工厂方法优缺点
简单工厂模式和工厂方法模式看起来很相似，本质区别就在于，如果在包子店中直接创建包子产品，是依赖具体包子店的，扩展性、弹性、可维护性都较差，而如果将实例化的代码抽象出来，不再依赖具体包子店，而是依赖于抽象的包子接口，使对象的实现从使用中解耦，这样就拥有很强的扩展性了，也可以称为 『依赖倒置原则』
工厂方法模式的优缺点


优点:


* 符合“开闭”原则，具有很强的的扩展性、弹性和可维护性。修改时只需要添加对应的工厂类即可

* 使用了依赖倒置原则，依赖抽象而不是具体，使用（客户）和实现（具体类）松耦合

* 客户只需要知道所需产品的具体工厂，而无须知道具体工厂的创建产品的过程，甚至不需要知道具体产品的类名。

缺点:

* 每增加一个产品时，都需要一个具体类和一个具体创建者，使得类的个数成倍增加，导致系统类数目过多，复杂性增加

* 对简单工厂，增加功能修改的是工厂类；对工厂方法，增加功能修改的是产品类。

