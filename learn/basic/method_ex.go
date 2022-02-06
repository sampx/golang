package basic

import (
	"fmt"
)

const (
	//WHITE 白色
	WHITE = iota
	//BLACK 黑色
	BLACK
	//BLUE 蓝色
	BLUE
	//RED 红色
	RED
	//YELLOW 黄色
	YELLOW
)

//Color 颜色类
type Color byte

//Box 盒子类
type Box struct {
	width, height, depth float64
	color                Color
}

//BoxList Class
type BoxList []Box //a slice of boxes

//Volume 容积
func (b *Box) Volume() float64 {
	return b.width * b.height * b.depth
}

//SetColor 设置盒子颜色
func (b *Box) SetColor(c Color) {
	b.color = c
}

//BiggestColor 获得最大盒子的颜色
func (bl BoxList) BiggestColor() *Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return &k
}

//PaintItBlack 设置所有盒子为黑色
func (bl BoxList) PaintItBlack() { //BoxList slice本身是指针
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

//name 获得颜色的名称
func (c *Color) name() string {
	names := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return names[*c]
}

//Rectangle class
type Rectangle struct {
	width, height float64
}

//Circle class
type Circle struct {
	radius float64
}

func MethodComplexTest() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.name())
	fmt.Println("The biggest one is", boxes.BiggestColor().name())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.name())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().name())
}

func CircleAreaTest() {
	var c1 Circle
	c1.radius = 5.00
	p := &c1
	fmt.Println("Area of Circle(c1) = ", p.area())
}

func RectangleAreaTest() {
	r1 := Rectangle{width: 3.2, height: 1.8}
	fmt.Println("Area of Rectangle(r1) = ", (&r1).area())
}

//该 method 属于 Circle 类型对象中的方法
func (c *Circle) area() float64 { //func后面增加了一个receiver(也就是method所依从的主体)
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
