# switch

## 代码实例

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//case语句中可以使用多个表达式，用逗号分隔
	//default为默认分支，当case语句匹配不到后使用
	i := 11
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 10:
		fmt.Println("ten")
	default:
		fmt.Println(i)
	}
	// case后面跟随多个表达式的实例
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("today is weekend")

	default:
		fmt.Println("today is workday")
	}

	//switch可以不带表达式，是if/else的另一个表现形式
	num := 10
	switch {
	case num < 10:
		fmt.Printf("%d < 5\n", num)
	case num > 10:
		fmt.Printf("%d>10\n", num)
	default:
		fmt.Printf("%d = 10\n", num)
	}

}
```

## 执行结果

```bash
write 11 as 11
today is workday
10 = 10
```
