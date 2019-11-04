# 常量

## 程序实例

```go
package main

import (
	"fmt"
	"math"
)

const s string = "hello const"

func main() {
	// go 支持字符、字符串、布尔、数值常量
	// 关键子const
	// const可以出现在任何var语句可以出现的地方
	// 常数表达式可以执行任意精度的运算
	// 数值常量没有确定的类型，直到它被给定了一个类型，如一次显示的类型转化

	fmt.Printf("%T\n", s) //string

	const n= 1000
	fmt.Printf("%T\n", n) //默认类型  int
	fmt.Println(n)

	const n1= 1000 / 2.0
	fmt.Printf("%T\t\n", n1) //转换为新的类型 float
	fmt.Println(n1)

	//math Sin函数需要一个float64的参数
	fmt.Println(math.Sin(n))   //n从常量默认int类型，变成float64类型

}

```

## 实例结果

```bash
string
int
1000
float64	
500
0.8268795405320026

```