
```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
```

Объяснение: defer - это отложенный вызов функции. Все отложенные вызовы складываются в стек и вызываются по его правилам,
то есть в обратном порядке LIFO(LAST IN FIRST OUT). Так как у функции test есть именованный параметр, то вызов defer может его
изменить после выхода функции.