# 函数

- 不定参函数
- 多返回值示例

# 示例程序

```go
package main

import "fmt"

//a func example
func plus(a, b int) (error, int) { //plus 是函数名称 ； a b是函数参数； error int是函数返回值
	return nil, a + b
}

//不定参函数示例
func sum1(sums ...int) int {
	total := 0
	for _, num := range sums {
		total += num
	}

	return total

}

func main() {
	err, sum := plus(4, 4) //plus 值需要用变量接一下，如果想丢弃返回值使用_接收
	if err == nil {
	}
	fmt.Println(sum)

	sum2 := sum1(1, 2, 3, 4, 5)
	fmt.Println(sum2)

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(sum1(s1...)) //使用s1...作为参数传给不定参函数

}


```

## 执行结果

```bash
8
15
45
```