# 匿名函数和闭包

- 匿名函数：没有名字的函数

- 闭包： 附函数可以访问主函数的变量 匿名函数都是闭包

## 函数示例

```go
package main

import "fmt"

func main() {
	var num int
	num = 9
	fmt.Println("1. 变量num 的原始值:", num)
	//匿名函数的用法一：
	func() {
		num++
		fmt.Println("2. num in 匿名函数:", num)
	}()
	//匿名函数的用法二
	f := func() {
		num++
		fmt.Println("3. num in 匿名函数2：", num)
	}
	f()

	f2 := func() {
		num++
		fmt.Println("4. use func type print num value:", num)
	}
	//f2() 注意这里没有对匿名函数2进行调用
	//使用type定义函数类型的变量，使用变量传参调用匿名函数
	type FuncType func()
	var f1 FuncType
	f1 = f2
	f1()

	fmt.Println("Last. num in 主函数:", num)

	//匿名函数，带参数，带返回值示例
	f3 := func(a, b int) int {
		return a + b
	}
	sum := f3(6, 6)
	fmt.Println("5.", sum)

	aa := test()
	fmt.Println("6. ", aa())
	fmt.Println("6. ", aa())
	fmt.Println("6. ", aa())

}

//闭包示例，通过闭包函数实现调用累加

func test() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

```

## 执行结果

```bash
1. 变量num 的原始值: 9
2. num in 匿名函数: 10
3. num in 匿名函数2： 11
4. use func type print num value: 12
Last. num in 主函数: 12
5. 12
6.  1
6.  2
6.  3
```
