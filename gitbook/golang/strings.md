# 字符串处理

- strings包常用操作实例

```go
package main

import (
	"fmt"
	"strings"
)

//strings 包常用操作
func main() {
	//定义变量
	str1 := "hello,world"
	str2 := "h"
	s1 := []string{"123", "456", "789"}

	//Contains  Contains(str,substr) bool
	fmt.Println(strings.Contains(str1, str2)) //true

	//Join  连接字符串 参数1 切片，参数2 链接符  返回值 string
	fmt.Println(strings.Join(s1, ""))  //123456789
	fmt.Println(strings.Join(s1, "-")) //123-456-789

	//Index 找子字符串在字符串中的位置，如果找不到返回-1
	fmt.Println(strings.Index(str1, str2))

	//Repeat 字符串重复打印n次  参数一，字符串 参数二 打印的次数
	fmt.Println(strings.Repeat(str2, 2))

	//Split 字符串分割  返回值是slice
	fmt.Println(strings.Split(str1, ",")) //[hello world]

	//ToUpper 小写字母转成大写
	fmt.Println(strings.ToUpper(str1))

	//TrimSpace 去除字符串两端的空格
	fmt.Println(strings.TrimSpace("   xxx hehe"))

	// Replace 替换关键字  	//n 代表替换次数，如果小于0 代表全部替换
	s := "性感在线取名，性感" //  屏蔽敏感词
	fmt.Println(strings.Replace(s, "性感", "**", -1))

	//字符串比较 3方法  == Compare EqualFold
	fmt.Println(str1 == str2) //false
	//Compare 0 相等 1不等
	fmt.Println(strings.Compare(str1, str2))
	//EqualFold utf-8小写编码 是否相等
	fmt.Println(strings.EqualFold("go", "gO")) //true
}



```