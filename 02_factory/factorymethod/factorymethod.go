package factorymethod

// 添加工厂接口（包子店接口）
type BunShopInterface interface {
	Generate(t string) Bun
}

// 添加产品接口（包子接口）
type Bun interface {
	create() string
}

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

func newBunShopFactory(s string) BunShopInterface {
	switch s {
	case "QS":
		return &QSBunShop{}
	case "GD":
		return &GDBunShop{}
	default:
		return nil
	}
}

func Run(bunShop, bunType string) string {
	f := newBunShopFactory(bunShop)
	b := f.Generate(bunType)
	bun := b.create()
	return bun
}
