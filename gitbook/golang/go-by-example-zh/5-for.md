# for 

## 程序实例

```go
package main

import "fmt"

func main() {
	//for 是Go中唯一的循环结构。
	//for 单个循环条件
	var i int //零值为0
	for i < 3 {
		fmt.Printf("%d\t", i)
		i++
	}

	fmt.Printf("\n")
	//for经典循环 初始化:条件:后续形式
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\t", i)
	}

	fmt.Printf("\n")
	//不带条件的for循环将一直执行，直到遇到break或return跳出循环
	for {
		fmt.Println("hello")
		break
	}
}

```

## 执行结果

```bash
0	1	2	
0	1	2	3	4	5	6	7	8	9	10	
hello
```

## 备注

for还有其他使用形式，后面会遇到： 如for range  、 channels等

