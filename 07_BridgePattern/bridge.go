package bridgepattern

type DrawApi interface {
	DrwaCircle(radius, x, y int)
}

//ShapeCircle 桥接模式的实体类
type ShapeCircle struct {
	Radius  int
	X       int
	Y       int
	drawAPI DrawApi
}

//NewShapeCircle 实例化桥接模式实体类
func NewShapeCircle(radius, x, y int, drawAPI DrawApi) *ShapeCircle {
	return &ShapeCircle{
		Radius:  radius,
		X:       x,
		Y:       y,
		drawAPI: drawAPI,
	}
}

//Draw 实体类的Draw方法
func (sc *ShapeCircle) Draw() {
	sc.drawAPI.DrwaCircle(sc.Radius, sc.X, sc.Y)
}
