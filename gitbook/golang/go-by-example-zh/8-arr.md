# array
 
 - 在Go中，数组是一个**固定长度**长度数列
 - 元素的**类型和长度**都是数组类型的一部分
 - 一个数组中，每个元素的类型必须是一样的

## 代码实例

```go
package main

import "fmt"

func main() {

	var arr1 [3]string // 声明一个数组 注意类型和长度
	var arr3 [3]int
	fmt.Println("字符串数组的零值为: ", arr1) //打印0值
	fmt.Println("int数组的零值为：", arr3)
	fmt.Println("数组长度为: ", len(arr1)) //数组长度
	fmt.Println("数组容量为: ", cap(arr1)) //数组容量
	//数组赋值 array[index] = value
	arr1[0] = "Hello"
	arr1[1] = "Golang"
	arr1[2] = "World!"
	fmt.Println(arr1)

	//数组初始化同时赋值
	arr2 := [3]string{"Ni", "Hao", "Golang"}
	fmt.Println(arr2)
	//数组遍历，for range
	for k, v := range arr1 { // for range 后面有介绍
		fmt.Printf("index: %d\t value:%s\n", k, v)
	}
	//数组遍历 for循环
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("index2 is %d\t value2 is %s\n", i, arr2[i])
	}

	//多维数组实例
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
```

## 运行结果

```bash
字符串数组的零值为:  [  ]
int数组的零值为： [0 0 0]
数组长度为:  3
数组容量为:  3
[Hello Golang World!]
[Ni Hao Golang]
index: 0	 value:Hello
index: 1	 value:Golang
index: 2	 value:World!
index2 is 0	 value2 is Ni
index2 is 1	 value2 is Hao
index2 is 2	 value2 is Golang
2d:  [[0 1 2] [1 2 3]]
```

## 备注

- 数组由于长度不可变，使用会有较多局限性，在Go程序中，slice会使用的更多，见后文。