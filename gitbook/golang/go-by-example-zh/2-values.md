# Values

## 代码实例

```go
package main

import "fmt"

func main() {
	fmt.Println("hello" + " " + "world") //字符串通过加号 + 链接
	fmt.Println("5*5=", 5*5)             //数学运算,整型运算  25
	fmt.Println("7.0/3.0", 7.0/3.0)      //浮点运算 2.3333333333333335
	fmt.Println(true && false)           //逻辑运算 与 false
	fmt.Println(true || false)           //逻辑运算 或 true
	fmt.Println(!true)                   //逻辑运算 非 false
}
```

## 实例结果

```bash
go run 2-values.go
hello world
5*5= 25
7.0/3.0 2.3333333333333335
false
true
false
```

