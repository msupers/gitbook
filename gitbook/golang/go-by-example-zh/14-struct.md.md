# struct

- 结构体是一组字段的集合，属性可以各不相同

## 程序实例
```go

package main

import "fmt"

func main() {

	// test1 定义结构体
	type struct1 struct {
		Age     int
		Name    string
		Address string
	}
	// test 1 打印zero values
	var s1 struct1
	fmt.Println("1. 结构体空值 ", s1)

	fmt.Println("2. 结构体赋值", struct1{1, "zhangsan", "beijing"})
	// 取 struct元素值
	s2 := struct1{1, "lisi", "shanghai"}
	fmt.Println("3.获取结构体元素", s2.Name)
	//使用结构体指针获取元素值
	p1 := &s2 //生成一个结构体指针
	fmt.Println("4.通过指针获取元素", p1.Name)
}
```

## 执行结果

```bash
1. 结构体空值  {0  }
2. 结构体赋值 {1 zhangsan beijing}
3.获取结构体元素 lisi
4.通过指针获取元素 lisi
```

