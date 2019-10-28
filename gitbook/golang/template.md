# text/template

Package template implements data-driven templates for generating textual output.


## 模板使用流程

- 定义模板>解析模板>数据驱动模板

```go
package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	type UserInfo struct {
		Name string
		Age  int
	}
//初始化结构体
	var user1 UserInfo
	user1.Name = "ZhangSan"
	user1.Age = 22
// New创建模板对象，Parse解析模板字符串，Execute数据驱动模板，对标识位进行替换
	tpl, err := template.New("test").Parse("my name is {{.Name}},my age is {{.Age}}")
	if err != nil {
		fmt.Println(err)
	}
	tpl.Execute(os.Stdout, user1)
}

```
执行结果如下：
my name is ZhangSan,my age is 22

## Actions

### 注释

- 说明：执行时会忽略。可以多行。注释不能嵌套，并且必须紧贴分界符始止


- 语法：{{/*comment*/}}

- 示例：

```go
package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	Name := "Tom"
	str := "My Name is {{.}}{{/*This is a comment*/}}"
	tpl,err := template.New("test").Parse(str)
	if err != nil{fmt.Println(err)}
	tpl.Execute(os.Stdout,Name)
}

```
执行结果： My Name is Tom
