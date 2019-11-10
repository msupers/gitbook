# pointer

## 示例程序

```go
package main

import "fmt"

func varExample(varal int) {
	varal = 0
	//fmt.Println(varal)
}

func pointerExample(ptr *int) {
	*ptr = 0
	//fmt.Println(*ptr)
}

func main() {
	// fmt.Printf 打印变量的地址 %v 或 %p
	var i int = 100
	fmt.Printf("1. var i address in MEM %p\n", &i)
	fmt.Printf("2. var i address in MEM %v\n", &i)

	varExample(i)      //varExample 不可以改变main函数中的i值
	fmt.Println(i)

	pointerExample(&i)  //poomterExample 可以改变main函数中的值 因为变量在内存引用是同一个地址
	fmt.Println(i)

}


```

## 执行结果

```bash
1. var i address in MEM 0xc00000a0b8
2. var i address in MEM 0xc00000a0b8
100
0
```