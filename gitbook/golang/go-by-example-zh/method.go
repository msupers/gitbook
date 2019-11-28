// - go 支持在结构体中定义方法
// - 可以使用值类型或者指针类型的**接收器**定义方法
// - 可以使用指针调用方法，避免在调用时候产生一个拷贝

package main

import "fmt"

type rect struct { //rect 矩形
	width  int
	height int
}

func (r rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {

	return 2*r.height + 2*r.width
}
func main() {
	//r := rect{3, 2}
	////r.area(4,3)
	//fmt.Println(r.area())

	var r rect
	r.width = 4
	r.height = 5

	rp := &r
	fmt.Println("面积是", r.area())
	fmt.Println("周长是", rp.perim())
}
