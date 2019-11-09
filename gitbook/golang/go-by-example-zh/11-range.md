# for range

- go语言可以使用for + range 迭代各种数据结构

## 代码示例

```go
package main

import (
	"fmt"
)

func main() {

	// example of for range迭代slice
	var s1 = []string{"hello", "golang", "world", "thanks", "you"}
	for k, v := range s1 {
		fmt.Printf("slice 下标: %d, 值: %s\n", k, v)
	}

	// 不需要索引示例,使用_丢弃索引
	for _, v := range s1 {
		fmt.Printf("s1元素依次为: %s\t", v)
		fmt.Println("") //换行
	}

	//使用for range遍历map
	map1 := map[int]string{1: "hello", 2: "golang", 3: "world"}
	for k, v := range map1 {
		fmt.Printf("map key: %d value:%s\n", k, v)
	}

	//复习使用make创建map
	map2 := make(map[string]string)
	map2["a"] = "hello"
	map2["b"] = "zhang san"
	for _, v := range map2 {
		fmt.Println("map2 have value：", v)
	}

	//for range迭代Unicode 编码
	string1 := "Hello"
	for i, c := range string1 {
		fmt.Printf("key: %d  type: %T \t value: %d  value to string %s\n", i, c, c, string(c))
	}

}
```

## 执行结果

```bash
slice 下标: 0, 值: hello
slice 下标: 1, 值: golang
slice 下标: 2, 值: world
slice 下标: 3, 值: thanks
slice 下标: 4, 值: you
s1元素依次为: hello	
s1元素依次为: golang	
s1元素依次为: world	
s1元素依次为: thanks	
s1元素依次为: you	
map key: 1 value:hello
map key: 2 value:golang
map key: 3 value:world
map2 have value： hello
map2 have value： zhang san
key: 0  type: int32 	 value: 72  value to string H
key: 1  type: int32 	 value: 101  value to string e
key: 2  type: int32 	 value: 108  value to string l
key: 3  type: int32 	 value: 108  value to string l
key: 4  type: int32 	 value: 111  value to string o
```


