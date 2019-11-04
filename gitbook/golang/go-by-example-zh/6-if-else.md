# if-else

## 程序实例

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	//可以只有if 没有else
	//if 后面是一个bool类型的值或表达式
	// if 条件语句之前可以有一个语句，在此声明的变量，在后面的所有条件分支都可以使用。
	// if else if  else 可以一起使用

	if true {
		fmt.Println("if")
	}

	if 7%2 == 0 {
		fmt.Println("7 is 偶数")
	} else {
		fmt.Println("7 is 奇数")
	}
	// strings.Contains 参考 go标准库 strings
	if firstName, lastName := "zhang", "san"; strings.Contains(firstName, "zh") {
		fmt.Printf("first name is： %s, contains str zh. Full Name is %s\t%s\n", firstName, firstName, lastName)
	} else {
		fmt.Printf("first name is： %s, not contains str zh. Full name is %s\t%s\n", firstName, firstName, lastName)
	}

	// if {} else if {} else {} example
	if num := 9; num < 0 {
		fmt.Printf("%d is 负数\n", num)
	} else if num > 0 {
		fmt.Printf("%d is 正数\n", num)
	} else {
		fmt.Println("%d is 0\n", num)
	}
}

```

## 执行结果

```bash
go run 6-if-else.go
if
7 is 奇数
first name is： zhang, contains str zh. Full Name is zhang	san
9 is 正数
```
