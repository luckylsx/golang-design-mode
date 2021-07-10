package builderpattern

import "fmt"

type Builder interface {
	NewProduct()            // 创建一个空产品
	BuildWheels()           // 建造轮子
	BuildChassis()          // 建造底盘
	BuildSeat()             // 建造驾驶位
	GetResult() interface{} // 获取建造好的产品
}

var (
	wheels  = "wheels"
	chassis = "chassis"
	seat    = "seat"
)

type Car struct {
	Wheels  string
	Chassis string
	Seat    string
}

type CarBuilder struct {
	Car *Car
}

func (c *CarBuilder) NewProduct() {
	c.Car = new(Car)
}

func (c *CarBuilder) BuildWheels() {
	c.Car.Wheels = wheels
}

func (c *CarBuilder) BuildChassis() {
	c.Car.Chassis = chassis
}

func (c *CarBuilder) BuildSeat() {
	c.Car.Seat = seat
}

func (c *CarBuilder) GetResult() interface{} {
	return c.Car
}

func (c *Car) Show() string {
	return fmt.Sprintf("%s %s %s", wheels, chassis, seat)
}

type Director struct {
	Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.Builder = builder
}

func (d *Director) Generate() *Car {
	d.Builder.NewProduct()
	d.Builder.BuildWheels()
	d.Builder.BuildChassis()
	d.Builder.BuildSeat()
	return d.Builder.GetResult().(*Car)
}
