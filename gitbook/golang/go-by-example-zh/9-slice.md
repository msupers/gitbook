# slice

- slice 类型又它包含的元素类型决定，类型不包括元素个数。
- silce内建方法 len() append() cap()
- slice可以被copy

## 程序实例

```go
package main

import "fmt"

func main() {

	//声明一个slice
	s1 := make([]string, 3)
	fmt.Println(s1)      //打印slice 0值
	fmt.Println(len(s1)) //打印slice 长度
	fmt.Println(cap(s1)) //打印slice 容量

	//append示例
	s1 = append(s1, "4") //append后，原来长度为3的slice长度变成4
	fmt.Println(len(s1))

	//slice copy示例

	var s2 []string
	s2 = append(s2, "1", "2", "3", "4")

	c := make([]string, len(s2)) // var c []string这种不定长的不可以copy
	copy(c, s2)
	fmt.Println("slice c copy from s2:", c)

	//对切片进行切片 [)区间    slice[low:high]
	//取0-3下标的元素
	fmt.Println(s2[0:4])      //没有下标为4的元素，这样写是因为冒号后面是开区间
	fmt.Println(s2[1:3])      //子集
	fmt.Println(s2[:])        //所有元素
	fmt.Println(s2[:len(s2)]) //len值相等于最大下标号+1

}


```

## 运行结果

```bash
[  ]
3
3
4
slice c copy from s2: [1 2 3 4]
[1 2 3 4]
[2 3]
[1 2 3 4]
[1 2 3 4]
```
