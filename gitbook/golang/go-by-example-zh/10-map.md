# map


- map 是go的内置关联数据类型 其他语言也称**哈希**或**字典**
- 当从map中取值，使用**第二个**返回值判断是否存在于map中，不存在的话反馈bool类型的false，存在返回true
- map中的key value对没有固定位置，一般都从key获取value，不能从value获取key

## 代码实例 

```go
package main

import "fmt"

func main() {
	//创建一个空map  make(map[key-type]val-type)
	map1 := make(map[int]string)
	fmt.Println(map1) //打印空值 map[]

	// map初始化并赋值
	map3 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(map3)

	// map 赋值
	map1[1] = "Hello"
	map1[3] = "Go Map"
	map1[2] = "abc"
	fmt.Println(map1)    // map[1:Hello 2:abc 3:Go Map]
	fmt.Println(map1[2]) //取map中某一个特定值，name[key] 打印对应的value   abc
	//fmt.Println(cap(map1))  map 没有cap函数

	//打印map的长度 len()
	fmt.Println(len(map1)) //3

	//从map中删除一个值
	delete(map1, 2)
	fmt.Println(map1) //map[1:Hello 3:Go Map]
	// 判断map中是否存在某个元素
	_, m := map1[2]
	fmt.Printf("%T\n", m)
	if m {
		fmt.Println("map1 中存在 ", map1[2])
	} else {
		fmt.Println("map1中不存在", map1[2])
	}

}
```

## 运行结果

```bash
map[]
map[one:1 three:3 two:2]
map[1:Hello 2:abc 3:Go Map]
abc
3
map[1:Hello 3:Go Map]
bool
map1中不存在
```
