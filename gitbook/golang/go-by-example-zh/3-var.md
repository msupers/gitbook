# 变量

## 程序实例

```go
package main

import "fmt"

func main() {

	// 1. var 声明1个或多个变量
	// 2. go 自动判断以及初始化的变量类型
	// 3. 声明变量没有赋值时，变量将会初始化未零值 如: int的零值时是0 bool的零值是false
	// 4. := 语句是申明并初始化变量的简写，f := "hello" 相当于 var f string = "hello"

	var str1 string = "i am str1"
	str2 := "i am str2"

	var str3 string
	str3 = "i am str3"

	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
	fmt.Println("str3:", str3)

	//一次定义多个变量
	var b, c = 1, 3
	fmt.Println("b=", b)
	fmt.Println("c=", c)

	//打印零值
	var n1 int
	var n2 string
	var n3 []string
	var n4 []int
	var n5 [5]int
	var n6 map[int]string

	fmt.Println("int的零值是:", n1)
	fmt.Println("string的零值是:", n2)
	fmt.Println("[]string的零值是", n3)
	fmt.Println("[]int的零值是", n4)
	fmt.Println("[5]int的零值是", n5)
	fmt.Println("map[int]string的零值是", n6)

}
```


## 实例结果

```bash
str1: i am str1
str2: i am str2
str3: i am str3
b= 1
c= 3
int的零值是: 0
string的零值是: 
[]string的零值是 []
[]int的零值是 []
[5]int的零值是 [0 0 0 0 0]
map[int]string的零值是 map[]

```